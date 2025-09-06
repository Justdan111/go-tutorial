package handlers

import (
	"net/http"
	"encoding/json"
	"notes-api/internal/models"
	"notes-api/internal/auth"
	"notes-api/internal/storage"

	"github.com/google/uuid"
	
)

type AuthHandler struct {
	storage  *storage.MemoryStorage
	JWTService  *auth.JWTService
	passwordService *auth.PasswordService

}

func NewAuthHandler(storage *storage.MemoryStorage, jwtService *auth.JWTService, passwordService *auth.PasswordService) *AuthHandler {
	return &AuthHandler{
		storage:  storage,
		JWTService: jwtService,
		passwordService: passwordService,
	}
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var req models.SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := models.ValidateSignupRequest(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}