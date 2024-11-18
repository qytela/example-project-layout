package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/qytela/example-project-layout/internal/app/api/handlers"
	"github.com/qytela/example-project-layout/internal/app/api/middleware"
)

func AuthRoutes(g *echo.Group, handler *handlers.AuthHandler) {
	route := "/auth"

	router := g.Group(route)

	routerProtect := g.Group(route)
	routerProtect.Use(middleware.Auth)

	routerRefreshProtect := g.Group(route)
	routerRefreshProtect.Use(middleware.AuthRefresh)

	router.POST("/login", handler.SignInWithEmailPassword)

	// Refresh token route
	routerRefreshProtect.POST("/refresh-token", handler.GenerateNewRefreshToken)

	// Protected route
	routerProtect.GET("/me", handler.GetUser)
}

func NoteRoutes(g *echo.Group, handler *handlers.NoteHandler) {
	router := g.Group("/notes")
	router.Use(middleware.Auth)

	router.GET("/user-notes", handler.GetUserNotes)
	router.GET("", handler.GetNotes)
	router.POST("", handler.StoreNote)
	router.PUT("/:id", handler.UpdateNote)
	router.DELETE("/:id", handler.DeleteNote)
}
