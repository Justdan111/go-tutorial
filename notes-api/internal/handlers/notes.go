package handlers

import (
	"encoding/json"
	"net/http"
	
	"notes-api/internal/storage"
	"strings"

	
)

type NotesHandler struct {
   storage *storage.MemoryStorage
}

func NewNotesHandler(storage *storage.MemoryStorage) *NotesHandler {
    return &NotesHandler{storage: storage}
} 

// checks the HTTP method of the request
func (h *NotesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        h.handleGet(w, r)
    case http.MethodPost:
        h.handlePost(w, r)
    case http.MethodPut:
        h.handlePut(w, r)
    case http.MethodDelete:
        h.handleDelete(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// Handle GET requests to retrieve notes
func (h *NotesHandler) handleGet(w http.ResponseWriter, r *http.Request) {
    path := strings.TrimPrefix(r.URL.Path, "/api/notes")

    if path == "" || path == "/" {
        // Get notes by user id
        userID := r.Context().Value("userID").(string)
        notes, err := h.storage.GetNotesByUserID(r.Context(), userID)
        if err != nil {
            http.Error(w, "Failed to get notes", http.StatusInternalServerError)
            return
}

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(notes)
        return
    }

    // Get a single note by ID
    noteID := strings.TrimPrefix(path, "/")
    note, err := h.storage.GetNote(r.Context(), noteID)
    if err != nil {
        http.Error(w, "Note not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(note)
}


