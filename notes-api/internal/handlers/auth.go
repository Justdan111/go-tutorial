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