package handlers

import (
	"encoding/json"
	"net/http"
	"notes-api/internal/models"
	"notes-api/internal/storage"
	"notes-api/internal/auth"

	"github.com/google/uuid"
)

