package services

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/qytela/example-project-layout/internal/app/api/repository"
	"github.com/qytela/example-project-layout/internal/app/api/requests"
	"github.com/qytela/example-project-layout/internal/app/api/responses"
	"github.com/qytela/example-project-layout/internal/pkg/exception"
	"github.com/qytela/example-project-layout/internal/pkg/utils"
)

type AuthService struct {
	repository *repository.AuthRepository
}

func NewAuthService(repository *repository.AuthRepository) *AuthService {
	return &AuthService{
		repository: repository,
	}
}

func (s *AuthService) SignInWithEmailPassword(c echo.Context) (*responses.AuthSigninResponse, error) {
	req := new(requests.AuthLoginRequest)
	if err := utils.ValidateRequest(c, req); err != nil {
		return nil, exception.NewInvalidRequest(err)
	}

	user, err := s.repository.SignInWithEmailPassword(req)
	if err != nil {
		return nil, exception.NewUnauthorized()
	}

	userAuthGrant, err := s.repository.StoreAndGetUserAuthGrant(user.ID)
	if err != nil {
		return nil, exception.NewBadRequest()
	}

	return &responses.AuthSigninResponse{
		UserAuthGrant: userAuthGrant,
		User: &responses.AuthMeResponse{
			ID:          user.ID,
			Role:        user.Role,
			Email:       user.Email,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
			Phone:       user.Phone,
			BannedUntil: user.BannedUntil,
		},
	}, nil
}

func (s *AuthService) GenerateNewRefreshToken(c echo.Context) (*responses.AuthSigninResponse, error) {
	userId := c.Get("userId").(uuid.UUID)

	user, err := s.repository.GetUser(userId)
	if err != nil {
		return nil, exception.NewUnauthorized()
	}

	userAuthGrant, err := s.repository.GenerateNewRefreshToken(userId)
	if err != nil {
		return nil, exception.NewBadRequest()
	}

	return &responses.AuthSigninResponse{
		UserAuthGrant: userAuthGrant,
		User: &responses.AuthMeResponse{
			ID:          user.ID,
			Role:        user.Role,
			Email:       user.Email,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
			Phone:       user.Phone,
			BannedUntil: user.BannedUntil,
		},
	}, nil
}

func (s *AuthService) GetUser(c echo.Context) (*responses.AuthMeResponse, error) {
	userId := c.Get("userId").(uuid.UUID)

	data, err := s.repository.GetUser(userId)
	if err != nil {
		return nil, exception.NewUnauthorized()
	}

	return &responses.AuthMeResponse{
		ID:          data.ID,
		Role:        data.Role,
		Email:       data.Email,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		Phone:       data.Phone,
		BannedUntil: data.BannedUntil,
	}, nil
}
