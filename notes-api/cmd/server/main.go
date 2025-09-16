package main

import (
    "log"
    "net/http"
    "os"
    "strconv"
    "time"
    
    "notes-api/internal/auth"
    "notes-api/internal/handlers"
    "notes-api/internal/middleware"
    "notes-api/internal/storage"
    pkgMiddleware "notes-api/pkg/middleware"
)

type Config struct {
    Port                string
    JWTSecret          string
    JWTExpiry          time.Duration
    BcryptCost         int
    RateLimitRequests  int
    RateLimitWindow    time.Duration
    AllowedOrigins     []string
}

func main() {
    // Load configuration
    config := loadConfig()
    
    // Initialize services
    storage := storage.NewMemoryStorage()
    jwtService := auth.NewJWTService(config.JWTSecret, config.JWTExpiry)
    passwordService := auth.NewPasswordService(config.BcryptCost)
    
    // Initialize middleware
    authMiddleware := middleware.NewAuthMiddleware(jwtService)
    rateLimiter := middleware.NewRateLimiter(config.RateLimitRequests, config.RateLimitWindow)
    
    // Initialize handlers
    authHandler := handlers.NewAuthHandler(storage, jwtService, passwordService)
    notesHandler := handlers.NewNotesHandler(storage)
    
    // Setup routes
    mux := http.NewServeMux()
    
    // Public routes (no authentication required)
    mux.HandleFunc("/api/auth/signup", authHandler.Signup)
    mux.HandleFunc("/api/auth/login", authHandler.Login)
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })
    
    // Protected routes (authentication required)
    mux.Handle("/api/notes", authMiddleware.RequireAuth(notesHandler))
    mux.Handle("/api/notes/", authMiddleware.RequireAuth(notesHandler))
    
    // Apply global middleware chain
    handler := pkgMiddleware.Logging(
        middleware.CORS(config.AllowedOrigins)(
            rateLimiter.Middleware(mux),
        ),
    )
    
    // Start server
    log.Printf("Server starting on port %s", config.Port)
    log.Printf("JWT expiry: %v", config.JWTExpiry)
    log.Printf("Rate limit: %d requests per %v", config.RateLimitRequests, config.RateLimitWindow)
    log.Fatal(http.ListenAndServe(":"+config.Port, handler))
}

func loadConfig() Config {
    return Config{
        Port:              getEnv("PORT", "8080"),
        JWTSecret:         getEnv("JWT_SECRET", "your-super-secret-jwt-key-change-this-in-production"),
        JWTExpiry:         getDurationEnv("JWT_EXPIRY", 24*time.Hour),
        BcryptCost:        getIntEnv("BCRYPT_COST", 12),
        RateLimitRequests: getIntEnv("RATE_LIMIT_REQUESTS", 100),
        RateLimitWindow:   getDurationEnv("RATE_LIMIT_WINDOW", time.Minute),
        AllowedOrigins:    []string{"http://localhost:3000", "http://localhost:8080"},
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func getIntEnv(key string, defaultValue int) int {
    if value := os.Getenv(key); value != "" {
        if intValue, err := strconv.Atoi(value); err == nil {
            return intValue
        }
    }
    return defaultValue
}

func getDurationEnv(key string, defaultValue time.Duration) time.Duration {
    if value := os.Getenv(key); value != "" {
        if duration, err := time.ParseDuration(value); err == nil {
            return duration
        }
    }
    return defaultValue
}