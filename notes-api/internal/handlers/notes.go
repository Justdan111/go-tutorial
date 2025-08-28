package handlers

import (
    "encoding/json"
    "net/http"
    "notes-api/internal/models"
    "notes-api/internal/storage"
    "strings"

     "github.com/google/uuid"
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
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowedh)
    }
}