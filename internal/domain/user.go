package domain

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Name         string
	Email        string
	PasswordHash []byte
}

func CheckPassword(u *User, password string) bool {
	passwordBytes := []byte(password)
	err := bcrypt.CompareHashAndPassword(u.PasswordHash, passwordBytes)

	return err == nil
}

func CreateUser(name string, email string, password string) *User {
	id := uuid.New()
	createdAt := time.Now()
	updatedAt := time.Now()
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic("Erro ao gerar o hash da senha: " + err.Error())
	}

	return &User{
		ID:           id,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
	}
}
