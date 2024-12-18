// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package http

import (
	"github.com/qytela/example-project-layout/internal/app/api/handlers"
	"github.com/qytela/example-project-layout/internal/app/api/repository"
	"github.com/qytela/example-project-layout/internal/app/api/services"
	"github.com/qytela/example-project-layout/internal/app/providers"
)

// Injectors from injector.go:

func InitializeAuthHandler() *handlers.AuthHandler {
	db := providers.ProvideDB()
	client := providers.ProvideSupabase()
	authRepository := repository.NewAuthRepository(db, client)
	authService := services.NewAuthService(authRepository)
	authHandler := handlers.NewAuthHandler(authService)
	return authHandler
}

func InitializeNoteHandler() *handlers.NoteHandler {
	db := providers.ProvideDB()
	noteRepository := repository.NewNoteRepository(db)
	noteService := services.NewNoteService(noteRepository)
	noteHandler := handlers.NewNoteHandler(noteService)
	return noteHandler
}
