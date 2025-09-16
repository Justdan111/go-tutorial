package handlers

import (
	"encoding/json"
	"net/http"
	"notes-api/internal/auth"
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
