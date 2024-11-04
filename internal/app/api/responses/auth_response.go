package responses

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/qytela/example-project-layout/internal/app/api/models"
)

type AuthSigninResponse struct {
	UserAuthGrant models.UserAuthGrant `json:"grant"`
	User          *AuthMeResponse      `json:"user"`
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
