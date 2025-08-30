package models

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User  `json:"user"`
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func ValidateSignupRequest(req SignupRequest) error {
	// Validate email
	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email is required")
	}
	if !emailRegex.MatchString(req.Email) {
		return errors.New("invalid email format")
	}

	// Validate password
	if len(req.Password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	if len(req.Password) > 100 {
		return errors.New("password must be less than 100 characters")
	}
	return nil
}

func ValidateLoginRequest(req LoginRequest) error {
	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email is required")
	}
	if strings.TrimSpace(req.Password) == "" {
		return errors.New("password is required")
	}
	return nil
}