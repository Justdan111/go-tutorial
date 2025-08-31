package middleware


import (
	"context"
	"net/http"
	"notes-api/internal/auth"
	"strings"
)

type Authmiddleware struct {
	jwtService *auth.JWTService
}

func NewAuthMiddleware(jwtService *auth.JWTService) *Authmiddleware {
	return &Authmiddleware{
		jwtService: jwtService,
	}
}
