package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/qytela/example-project-layout/internal/app/api/services"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) SignInWithEmailPassword(c echo.Context) error {
	data, err := h.service.SignInWithEmailPassword(c)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"status": true,
		"data":   data,
	})
}

func (h *AuthHandler) GetUser(c echo.Context) error {
	data, err := h.service.GetUser(c)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"status": true,
		"data":   data,
	})
}
