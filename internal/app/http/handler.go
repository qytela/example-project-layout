package http

import (
	"github.com/qytela/example-project-layout/internal/app/api/handlers"
)

type Handler struct {
	*handlers.AuthHandler
	*handlers.NoteHandler
}

func NewHandler() *Handler {
	AuthHandler := InitializeAuthHandler()
	NoteHandler := InitializeNoteHandler()

	return &Handler{
		AuthHandler,
		NoteHandler,
	}
}
