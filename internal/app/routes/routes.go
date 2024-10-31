package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/qytela/example-project-layout/internal/app/api/handlers"
	"github.com/qytela/example-project-layout/internal/app/api/middleware"
)

func AuthRoutes(g *echo.Group, handler *handlers.AuthHandler) {
	router := g.Group("/auth")

	routerProtect := g.Group("/auth")
	routerProtect.Use(middleware.Auth)

	router.POST("/login", handler.SignInWithEmailPassword)

	// Protected route
	routerProtect.GET("/me", handler.GetUser)
}

func NoteRoutes(g *echo.Group, handler *handlers.NoteHandler) {
	router := g.Group("/notes")
	router.Use(middleware.Auth)

	router.GET("", handler.GetNotes)
	router.POST("", handler.StoreNote)
	router.PUT("/:id", handler.UpdateNote)
	router.DELETE("/:id", handler.DeleteNote)
}
