package auth

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct {
	cost int
}

func NewPasswordService(cost int) *PasswordService {
	return &PasswordService{cost: cost}
}

// HashPassword hashes the given plain-text password using bcrypt
func (p *PasswordService) HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), p.cost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// CheckPassword compares a hashed password with a plain-text password
func (p *PasswordService) CheckPassword(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}