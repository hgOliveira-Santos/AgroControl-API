package domain

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	Description string
	Value       float64
	Date        time.Time

	UserID uuid.UUID
	User   User
}
