package repository

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/qytela/example-project-layout/internal/app/api/models"
	"github.com/qytela/example-project-layout/internal/app/api/requests"
	"github.com/qytela/example-project-layout/internal/pkg/auth"
	"github.com/qytela/example-project-layout/internal/pkg/utils"
	"github.com/supabase-community/supabase-go"
)

type AuthRepository struct {
	db             *sqlx.DB
	supabaseClient *supabase.Client
}

func NewAuthRepository(db *sqlx.DB, supabaseClient *supabase.Client) *AuthRepository {
	return &AuthRepository{
		db:             db,
		supabaseClient: supabaseClient,
	}
}

func (r *AuthRepository) SignInWithEmailPassword(req *requests.AuthLoginRequest) (models.User, error) {
	row, err := r.GetUserByEmail(req.Email)
	if err != nil {
		return row, err
	}

	if err := utils.ComparePassword(row.EncryptedPassword, req.Password); err != nil {
		return row, err
	}

	return row, nil
}

func (r *AuthRepository) StoreAndGetUserAuthGrant(userId uuid.UUID) (models.UserAuthGrant, error) {
	var userAuthGrant models.UserAuthGrant
	var userGrant models.UserAuthGrant

	token, err := auth.GenerateAccessToken(userId)
	if err != nil {
		return userGrant, err
	}

	tokenExpiredAt := time.Now().Add(time.Minute * 5)
	refreshTokenExpiredAt := time.Now().Add(time.Minute * 60)

	// Revoked all token before create new token
	if err := r.RevokedUserAuthGrant(userId); err != nil {
		return userGrant, err
	}

	// Begin transaction
	tx := r.db.MustBegin()

	queryInsert := `
		INSERT INTO user_auth_grants
			(user_id, access_token, access_token_revoked, access_token_expired_at)
		VALUES
			($1, $2, $3, $4)
		RETURNING *
	`
	if err := tx.QueryRowx(queryInsert, userId, token, false, tokenExpiredAt).StructScan(&userAuthGrant); err != nil {
		return userGrant, err
	}

	refreshToken, err := auth.GenerateRefreshToken(userAuthGrant.ID)
	if err != nil {
		return userGrant, err
	}

	queryUpdate := `
		UPDATE user_auth_grants
		SET refresh_token = $1, refresh_token_expired_at = $2
		WHERE user_id = $3 AND id = $4
		RETURNING *
	`
	if err := tx.Get(&userGrant, queryUpdate, refreshToken, refreshTokenExpiredAt, userId, userAuthGrant.ID); err != nil {
		return userGrant, err
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		// Rollback transaction if error
		tx.Rollback()
		return userGrant, err
	}

	return userGrant, nil
}

func (r *AuthRepository) GenerateNewRefreshToken(req *requests.AuthRefreshRequest) (models.UserAuthGrant, error) {
	var userAuthGrant models.UserAuthGrant

	query := `SELECT * FROM user_auth_grants WHERE refresh_token = $1 AND refresh_token_revoked = false`
	if err := r.db.Get(&userAuthGrant, query, req.RefreshToken); err != nil {
		return userAuthGrant, err
	}

	if userAuthGrant.RefreshTokenExpiredAt.Time.Unix() < time.Now().Add(-60*time.Minute).Unix() {
		// Revoked all token if expired
		if err := r.RevokedUserAuthGrant(userAuthGrant.UserID); err != nil {
			return userAuthGrant, err
		}

		return userAuthGrant, errors.New("refresh token has expired")
	}

	userGrant, err := r.StoreAndGetUserAuthGrant(userAuthGrant.UserID)
	if err != nil {
		return userAuthGrant, err
	}

	if _, err := auth.ValidateRefreshToken(userGrant.RefreshToken.String); err != nil {
		return userAuthGrant, err
	}

	return userGrant, nil
}

func (r *AuthRepository) RevokedUserAuthGrant(userId uuid.UUID) error {
	query := `
		UPDATE user_auth_grants
		SET access_token_revoked = true, refresh_token_revoked = true
	`
	if _, err := r.db.Exec(query); err != nil {
		return err
	}

	return nil
}

func (r *AuthRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	query := `SELECT * FROM auth.users WHERE email = $1`
	if err := r.db.Get(&user, query, email); err != nil {
		return user, err
	}

	return user, nil
}

func (r *AuthRepository) GetUser(userId uuid.UUID) (models.User, error) {
	var user models.User

	query := `SELECT * FROM auth.users WHERE id = $1`
	if err := r.db.Get(&user, query, userId); err != nil {
		return user, err
	}

	return user, nil
}
