package models

import (
	"time"

	"github.com/google/uuid"
)

type UserNote struct {
	UserID    uuid.UUID  `json:"user_id" db:"user_id"`
	NoteID    int8       `json:"note_id" db:"note_id"`
	Note      string     `json:"note" db:"note"`
	Order     int        `json:"order" db:"order"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
