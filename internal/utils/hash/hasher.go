package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordHasher interface {
	Hash(password string) (string, error)
	ComparePasswordHash(password, hash string) bool
}

type BcryptHasher struct {
}

func NewBcryptHasher() *BcryptHasher {
	return &BcryptHasher{}
}

func (h *BcryptHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", nil
	}
	return string(bytes), nil
}

func (h *BcryptHasher) ComparePasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}