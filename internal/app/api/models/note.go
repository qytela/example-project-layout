package models

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID        int8       `json:"id" db:"id"`
	UserID    uuid.UUID  `json:"user_id" db:"user_id"`
	Note      string     `json:"note" db:"note"`
	Order     int        `json:"order" db:"order"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
