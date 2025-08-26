package storage

import (
    "context"
    "errors"
    "notes-api/internal/models"
    "sync"
    "time"
)

type MemoryStorage struct {
    notes map[string]*models.Note
    mu    sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
    return &MemoryStorage{
        notes: make(map[string]*models.Note),
    }
}

func (s *MemoryStorage) CreateNote(ctx context.Context, note *models.Note) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    note.CreatedAt = time.Now()
    note.UpdatedAt = time.Now()
    s.notes[note.ID] = note
    return nil
}

func (s *MemoryStorage) GetNote(ctx context.Context, id string) (*models.Note, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    note, exists := s.notes[id]
    if !exists {
        return nil, errors.New("note not found")
    }
    return note, nil
}

func (s *MemoryStorage) GetAllNotes(ctx context.Context) ([]*models.Note, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    notes := make([]*models.Note, 0, len(s.notes))
    for _, note := range s.notes {
        notes = append(notes, note)
    }
    return notes, nil
}

func (s *MemoryStorage) UpdateNote(ctx context.Context, id string, updated *models.Note) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    note, exists := s.notes[id]
    if !exists {
        return errors.New("note not found")
    }
    
    note.Title = updated.Title
    note.Content = updated.Content
    note.UpdatedAt = time.Now()
    return nil
}

func (s *MemoryStorage) DeleteNote(ctx context.Context, id string) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    if _, exists := s.notes[id]; !exists {
        return errors.New("note not found")
    }
    
    delete(s.notes, id)
    return nil
}