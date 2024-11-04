package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type UserAuthGrant struct {
	ID                    int8            `json:"id" db:"id"`
	UserID                uuid.UUID       `json:"user_id" db:"user_id"`
	AccessToken           string          `json:"access_token" db:"access_token"`
	RefreshToken          *sql.NullString `json:"refresh_token" db:"refresh_token"`
	AccessTokenRevoked    bool            `json:"access_token_revoked" db:"access_token_revoked"`
	RefreshTokenRevoked   bool            `json:"refresh_token_revoked" db:"refresh_token_revoked"`
	AccessTokenExpiredAt  *time.Time      `json:"access_token_expired_at" db:"access_token_expired_at"`
	RefreshTokenExpiredAt *sql.NullTime   `json:"refresh_token_expired_at" db:"refresh_token_expired_at"`
	CreatedAt             *time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt             *time.Time      `json:"updated_at" db:"updated_at"`
}
