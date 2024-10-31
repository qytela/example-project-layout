//go:build wireinject
// +build wireinject

package http

import (
	"github.com/google/wire"
	"github.com/qytela/example-project-layout/internal/app/api/handlers"
	"github.com/qytela/example-project-layout/internal/app/api/repository"
	"github.com/qytela/example-project-layout/internal/app/api/services"
	"github.com/qytela/example-project-layout/internal/app/providers"
)

func InitializeAuthHandler() *handlers.AuthHandler {
	wire.Build(providers.ProvideDB, providers.ProvideSupabase, repository.NewAuthRepository, services.NewAuthService, handlers.NewAuthHandler)

	return nil
}

func InitializeNoteHandler() *handlers.NoteHandler {
	wire.Build(providers.ProvideDB, repository.NewNoteRepository, services.NewNoteService, handlers.NewNoteHandler)

	return nil
}
