package models

import (
    "errors"
    "time"
    "strings"
)


type Note struct {
    ID        string    `json:"id"`
    UserID    string    `json:"user_id"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type CreateNoteRequest struct {
    Title   string `json:"title"`
    Content string `json:"content"`
}

type UpdateNoteRequest struct {
    Title   string `json:"title"`
    Content string `json:"content"`
}

func ValidateCreateNoteRequest(req CreateNoteRequest) error {
    if strings.TrimSpace(req.Title) == "" {
        return errors.New("title is required")
    }
    if len(req.Title) > 100 {
        return errors.New("title must be less than 100 charaters")
    }
    if strings.TrimSpace(req.Content) == "" {
        return errors.New("content is required")
    }
    if len(req.Content) > 5000 {
        return errors.New("content must be less than 5000 characters")
    }
    return nil
}

func ValidateUpdateNoteRequest(req UpdateNoteRequest) error {
    return ValidateCreateNoteRequest(CreateNoteRequest(req))
}
