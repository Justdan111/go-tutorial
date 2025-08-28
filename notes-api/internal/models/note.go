package models

import (
    "errors"
    "time"
    "strings"
)

type Note struct {
    ID        string    `json:"id"`
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

func ValidateCreatedRequest(req CreateNoteRequest) error {
    if strings.TrimSpace(req.Title) == "" {
        return errors.New("Title is required")
    }
    if len(req.Title) > 100 {
        return errors.New("title must be less than 100 charaters")
    }
    if strings.TrimSpace(req.Content) == "" {
        return errors.New("content is required")
    }
    return nil
}