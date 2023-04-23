package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          int64
	Name        string
	Description string
	UserID      uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductOwner struct {
	Product *Product
	User    *User
}
