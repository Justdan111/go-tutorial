package models

import "time"

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