package responses

import (
	"database/sql"

	"github.com/google/uuid"
)

type AuthSigninResponse struct {
	AccessToken string          `json:"access_token"`
	User        *AuthMeResponse `json:"user"`
}

type AuthMeResponse struct {
	ID          uuid.UUID       `json:"id"`
	Role        string          `json:"role"`
	Email       string          `json:"email"`
	CreatedAt   *sql.NullTime   `json:"created_at"`
	UpdatedAt   *sql.NullTime   `json:"updated_at"`
	Phone       *sql.NullString `json:"phone"`
	BannedUntil *sql.NullTime   `json:"banned_until"`
}
