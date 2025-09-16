package handlers

import (
	"encoding/json"
	"net/http"
	"notes-api/internal/auth"
	"notes-api/internal/middleware"
	"notes-api/internal/models"
	"notes-api/internal/storage"

	"github.com/google/uuid"
)

type AuthHandler struct {
    storage         *storage.MemoryStorage
    jwtService      *auth.JWTService
    passwordService *auth.PasswordService
}

func NewAuthHandler(storage *storage.MemoryStorage, jwtService *auth.JWTService, passwordService *auth.PasswordService) *AuthHandler {
    return &AuthHandler{
        storage:         storage,
        jwtService:      jwtService,
        passwordService: passwordService,
    }
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
    var req models.SignupRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    if err := models.ValidateSignupRequest(req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Hash password
    hashedPassword, err := h.passwordService.HashPassword(req.Password)
    if err != nil {
        http.Error(w, "Failed to process password", http.StatusInternalServerError)
        return
    }

    // Create user
    user := &models.User{
        ID:           uuid.New().String(),
        Email:        req.Email,
        Password: hashedPassword,
    }

    if err := h.storage.CreateUser(r.Context(), user); err != nil {
        if err.Error() == "email already exists" {
            http.Error(w, err.Error(), http.StatusConflict)
        } else {
            http.Error(w, "Failed to create user", http.StatusInternalServerError)
        }
        return
    }

    // Generate JWT token
    token, err := h.jwtService.GenerateToken(*user)
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    response := models.LoginResponse{
        Token: token,
        User:  *user,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    var req models.LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    if err := models.ValidateLoginRequest(req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Get user by email
    user, err := h.storage.GetUserByEmail(r.Context(), req.Email)
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    // Check password
    if !h.passwordService.CheckPassword(req.Password, user.Password) {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    // Generate JWT token
    token, err := h.jwtService.GenerateToken(*user)
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    response := models.LoginResponse{
        Token: token,
        User:  *user,
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}


    func (h *NotesHandler) handlePost(w http.ResponseWriter, r *http.Request) {
        userID := middleware.GetUserIDFromContext(r.Context())
        if userID == "" {
            http.Error(w, "User not authenticated", http.StatusUnauthorized)
            return
        }
    
        var req models.CreateNoteRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }
        
        if err := models.ValidateCreateNoteRequest(req); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        
        note := &models.Note{
            ID:      uuid.New().String(),
            UserID:  userID, // Associate note with authenticated user
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
    
    func (h *NotesHandler) handlePut(w http.ResponseWriter, r *http.Request) {
        userID := middleware.GetUserIDFromContext(r.Context())
        if userID == "" {
            http.Error(w, "User not authenticated", http.StatusUnauthorized)
            return
        }
    
        path := strings.TrimPrefix(r.URL.Path, "/api/notes/")
        if path == "" {
            http.Error(w, "Note ID required", http.StatusBadRequest)
            return
        }
        
        // Check if note exists and belongs to user
        existingNote, err := h.storage.GetNote(r.Context(), path)
        if err != nil {
            http.Error(w, "Note not found", http.StatusNotFound)
            return
        }
        
        if existingNote.UserID != userID {
            http.Error(w, "Access denied", http.StatusForbidden)
            return
        }
        
        var req models.UpdateNoteRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }
        
        if err := models.ValidateUpdateNoteRequest(req); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        
        updateNote := &models.Note{
            Title:   req.Title,
            Content: req.Content,
        }
        
        if err := h.storage.UpdateNote(r.Context(), path, updateNote); err != nil {
            http.Error(w, "Failed to update note", http.StatusInternalServerError)
            return
        }
        
        // Get updated note to return
        note, _ := h.storage.GetNote(r.Context(), path)
        
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(note)
    }
    
    func (h *NotesHandler) handleDelete(w http.ResponseWriter, r *http.Request) {
        userID := middleware.GetUserIDFromContext(r.Context())
        if userID == "" {
            http.Error(w, "User not authenticated", http.StatusUnauthorized)
            return
        }
    
        path := strings.TrimPrefix(r.URL.Path, "/api/notes/")
        if path == "" {
            http.Error(w, "Note ID required", http.StatusBadRequest)
            return
        }
        
        // Check if note exists and belongs to user
        existingNote, err := h.storage.GetNote(r.Context(), path)
        if err != nil {
            http.Error(w, "Note not found", http.StatusNotFound)
            return
        }
        
        if existingNote.UserID != userID {
            http.Error(w, "Access denied", http.StatusForbidden)
            return
        }
        
        if err := h.storage.DeleteNote(r.Context(), path); err != nil {
            http.Error(w, "Failed to delete note", http.StatusInternalServerError)
            return
        }
        
        w.WriteHeader(http.StatusNoContent)
    }