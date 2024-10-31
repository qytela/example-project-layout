package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/qytela/example-project-layout/internal/app/api/models"
	"github.com/qytela/example-project-layout/internal/app/api/requests"
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
