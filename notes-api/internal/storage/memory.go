package storage

import (
    "context"
    "errors"
    "notes-api/internal/models"
    "sync"
    "time"
)

type MemoryStorage struct {
    users map[string]*models.User
    notes map[string]*models.Note
    usersByEmail map[string]*models.User
    mu    sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
    return &MemoryStorage{
        users:        make(map[string]*models.User),
        notes:        make(map[string]*models.Note),
        usersByEmail: make(map[string]*models.User),
    }
}

// User methods
func (s *MemoryStorage) CreateUser(ctx context.Context, user *models.User) error {
    s.mu.Lock()
    defer s.mu.Unlock()
    
    // Check if email already exists
    if _, exists := s.usersByEmail[user.Email]; exists {
        return errors.New("email already exists")
    }
    
    user.CreatedAt = time.Now()
    s.users[user.ID] = user
    s.usersByEmail[user.Email] = user
    return nil
}

func (s *MemoryStorage) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    user, exists := s.usersByEmail[email]
    if !exists {
        return nil, errors.New("user not found")
    }
    return user, nil
}

func (s *MemoryStorage) GetUserByID(ctx context.Context, id string) (*models.User, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    user, exists := s.users[id]
    if !exists {
        return nil, errors.New("user not found")
    }
    return user, nil
}

// Note methods (updated with user ownership)
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

func (s *MemoryStorage) GetNotesByUserID(ctx context.Context, userID string) ([]*models.Note, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    
    var userNotes []*models.Note
    for _, note := range s.notes {
        if note.UserID == userID {
            userNotes = append(userNotes, note)
        }
    }
    return userNotes, nil
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