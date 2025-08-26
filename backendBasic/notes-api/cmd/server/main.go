package main

import (
    "log"
    "net/http"
    "os"
    
    "your-project/internal/handlers"
    "your-project/internal/storage"
    "your-project/pkg/middleware"
)

func main() {
    // Load configuration
    port := getEnv("PORT", "8080")
    
    // Initialize storage
    storage := storage.NewMemoryStorage()
    
    // Initialize handlers
    notesHandler := handlers.NewNotesHandler(storage)
    
    // Setup routes
    mux := http.NewServeMux()
    mux.Handle("/api/notes", notesHandler)
    mux.Handle("/api/notes/", notesHandler)
    
    // Add health check endpoint
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })
    
    // Apply middleware
    handler := middleware.Logging(middleware.CORS(mux))
    
    // Start server
    log.Printf("Server starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, handler))
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}