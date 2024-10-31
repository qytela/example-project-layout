package http

import (
	"github.com/labstack/echo/v4"
	"github.com/qytela/example-project-layout/internal/app/routes"
)

type Server struct {
}

func NewServer(e *echo.Echo, handler *Handler) {
	apiGroup := e.Group("/api/v1")

	routes.AuthRoutes(apiGroup, handler.AuthHandler)
	routes.NoteRoutes(apiGroup, handler.NoteHandler)
}
