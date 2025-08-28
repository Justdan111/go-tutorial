package main

import (
    "log"
    "net/http"
    "notes-api/internal/handlers"
    "notes-api/internal/storage"
    "notes-api/pkg/middleware"
    "os"
)

func main() {
    // Load Configuration from environment variables
    port := getEnv("PORT", "8080")

    // Initialize in-memory storage
    storage := storage.NewMemoryStorage()

    // Initialize handlers
    notesHandler := handlers.NewNotesHandler(storage)

    // Set up routes
    mux := http.NewServeMux()
    mux.Handle("/api/notes", notesHandler)
    mux.Handle("/api/notes/", notesHandler)

    // Add Heathcheck endpoint
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    }   )

    // Apply middleware
    handler := middleware.Logging(middleware.CORS(mux))

    // Start the server
    log.Printf("Server is running on port %d", port)
    log.Printf("Try: curl http://localhost:%s/health", port)
    log.Fatal(http.ListenAndServe(":"+port, handler))
}

    func getEnv(key, defaultValue string) string {
        if value, exists := os.LookupEnv(key); exists {
            return value
        }
        return defaultValue
        
    }
