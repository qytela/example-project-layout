package models

import (
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	InstanceID               uuid.UUID       `json:"instance_id" db:"instance_id"`
	ID                       uuid.UUID       `json:"id" db:"id"`
	Aud                      string          `json:"aud" db:"aud"`
	Role                     string          `json:"role" db:"role"`
	Email                    string          `json:"email" db:"email"`
	EncryptedPassword        string          `json:"encrypted_password" db:"encrypted_password"`
	EmailConfirmedAt         *sql.NullTime   `json:"email_confirmed_at" db:"email_confirmed_at"`
	InvitedAt                *sql.NullTime   `json:"invited_at" db:"invited_at"`
	ConfirmationToken        string          `json:"confirmation_token" db:"confirmation_token"`
	ConfirmationSentAt       *sql.NullTime   `json:"confirmation_sent_at" db:"confirmation_sent_at"`
	RecoveryToken            string          `json:"recovery_token" db:"recovery_token"`
	RecoverySentAt           *sql.NullTime   `json:"recovery_sent_at" db:"recovery_sent_at"`
	EmailChangeTokenNew      string          `json:"email_change_token_new" db:"email_change_token_new"`
	EmailChange              string          `json:"email_change" db:"email_change"`
	EmailChangeSentAt        *sql.NullTime   `json:"email_change_sent_at" db:"email_change_sent_at"`
	LastSignInAt             *sql.NullTime   `json:"last_sign_in_at" db:"last_sign_in_at"`
	RawAppMetaData           string          `json:"raw_app_meta_data" db:"raw_app_meta_data"`
	RawUserMetaData          string          `json:"raw_user_meta_data" db:"raw_user_meta_data"`
	IsSuperAdmin             *sql.NullBool   `json:"is_super_admin" db:"is_super_admin"`
	CreatedAt                *sql.NullTime   `json:"created_at" db:"created_at"`
	UpdatedAt                *sql.NullTime   `json:"updated_at" db:"updated_at"`
	Phone                    *sql.NullString `json:"phone" db:"phone"`
	PhoneConfirmedAt         *sql.NullTime   `json:"phone_confirmed_at" db:"phone_confirmed_at"`
	PhoneChange              string          `json:"phone_change" db:"phone_change"`
	PhoneChangeToken         string          `json:"phone_change_token" db:"phone_change_token"`
	PhoneChangeSentAt        *sql.NullTime   `json:"phone_change_sent_at" db:"phone_change_sent_at"`
	ConfirmedAt              *sql.NullTime   `json:"confirmed_at" db:"confirmed_at"`
	EmailChangeTokenCurrent  string          `json:"email_change_token_current" db:"email_change_token_current"`
	EmailChangeConfirmStatus int16           `json:"email_change_confirm_status" db:"email_change_confirm_status"`
	BannedUntil              *sql.NullTime   `json:"banned_until" db:"banned_until"`
	ReauthenticationToken    string          `json:"reauthentication_token" db:"reauthentication_token"`
	ReauthenticationSentAt   *sql.NullTime   `json:"reauthentication_sent_at" db:"reauthentication_sent_at"`
	IsSsoUser                bool            `json:"is_sso_user" db:"is_sso_user"`
	DeletedAt                *sql.NullTime   `json:"deleted_at" db:"deleted_at"`
	IsAnonymous              bool            `json:"is_anonymous" db:"is_anonymous"`
}
