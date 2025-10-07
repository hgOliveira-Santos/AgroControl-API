package domain

import (
	"github.com/google/uuid"
)

type Safra struct {
	ID          uuid.UUID
	Name        string
	Description string
}
