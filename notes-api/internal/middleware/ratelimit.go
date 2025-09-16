package middleware

import (
    "net/http"
    "sync"
    "time"
)

type RateLimiter struct {
    requests map[string]*clientInfo
    mu       sync.RWMutex
    limit    int
    window   time.Duration
}

type clientInfo struct {
    requests  int
    resetTime time.Time
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
    rl := &RateLimiter{
        requests: make(map[string]*clientInfo),
        limit:    limit,
        window:   window,
    }
    
    // Clean up expired entries periodically
    go rl.cleanup()
    
    return rl
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Use IP address as client identifier
        // In production, you might use user ID from JWT for authenticated requests
        clientID := r.RemoteAddr
        
        if rl.isRateLimited(clientID) {
            http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
            return
        }
        
        next.ServeHTTP(w, r)
    })
}

func (rl *RateLimiter) isRateLimited(clientID string) bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    
    now := time.Now()
    
    client, exists := rl.requests[clientID]
    if !exists || now.After(client.resetTime) {
        // New client or window has reset
        rl.requests[clientID] = &clientInfo{
            requests:  1,
            resetTime: now.Add(rl.window),
        }
        return false
    }
    
    if client.requests >= rl.limit {
        return true
    }
    
    client.requests++
    return false
}

func (rl *RateLimiter) cleanup() {
    ticker := time.NewTicker(rl.window)
    defer ticker.Stop()
    
    for range ticker.C {
        rl.mu.Lock()
        now := time.Now()
        for clientID, client := range rl.requests {
            if now.After(client.resetTime) {
                delete(rl.requests, clientID)
            }
        }
        rl.mu.Unlock()
    }
}