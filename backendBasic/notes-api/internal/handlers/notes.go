package handlers

import (
    "encoding/json"
    "net/http"
    "strings"
    
    "github.com/google/uuid"
    "your-project/internal/models"
    "your-project/internal/storage"
)

type NotesHandler struct {
    storage *storage.MemoryStorage
}

func NewNotesHandler(storage *storage.MemoryStorage) *NotesHandler {
    return &NotesHandler{storage: storage}
}

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

func (h *NotesHandler) handleGet(w http.ResponseWriter, r *http.Request) {
    path := strings.TrimPrefix(r.URL.Path, "/api/notes")
    
    if path == "" || path == "/" {
        // Get all notes
        notes, err := h.storage.GetAllNotes(r.Context())
        if err != nil {
            http.Error(w, "Failed to get notes", http.StatusInternalServerError)
            return
        }
        
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(notes)
        return
    }
    
    // Get specific note
    noteID := strings.TrimPrefix(path, "/")
    note, err := h.storage.GetNote(r.Context(), noteID)
    if err != nil {
        http.Error(w, "Note not found", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(note)
}

func (h *NotesHandler) handlePost(w http.ResponseWriter, r *http.Request) {
    var req models.CreateNoteRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    if err := h.validateCreateRequest(req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    note := &models.Note{
        ID:      uuid.New().String(),
        Title:   req.Title,
        Content: req.Content,
    }
    
    if err := h.storage.CreateNote(r.Context(), note); err != nil {
        http.Error(w, "Failed to create note", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(note)
}

func (h *NotesHandler) validateCreateRequest(req models.CreateNoteRequest) error {
    if strings.TrimSpace(req.Title) == "" {
        return errors.New("title is required")
    }
    if len(req.Title) > 100 {
        return errors.New("title must be less than 100 characters")
    }
    if strings.TrimSpace(req.Content) == "" {
        return errors.New("content is required")
    }
    return nil
}