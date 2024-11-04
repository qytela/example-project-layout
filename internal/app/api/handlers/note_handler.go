package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/qytela/example-project-layout/internal/app/api/services"
)

type NoteHandler struct {
	service *services.NoteService
}

func NewNoteHandler(service *services.NoteService) *NoteHandler {
	return &NoteHandler{
		service: service,
	}
}

func (h *NoteHandler) GetUserNotes(c echo.Context) error {
	data, err := h.service.GetUserNotes(c)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"status": true,
		"data":   data,
	})
}

func (h *NoteHandler) GetNotes(c echo.Context) error {
	data, err := h.service.GetNotes(c)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"status": true,
		"data":   data,
	})
}

func (h *NoteHandler) StoreNote(c echo.Context) error {
	data, err := h.service.StoreNote(c)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"status": true,
		"data":   data,
	})
}

func (h *NoteHandler) UpdateNote(c echo.Context) error {
	data, err := h.service.UpdateNote(c)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"status": true,
		"data":   data,
	})
}

func (h *NoteHandler) DeleteNote(c echo.Context) error {
	if err := h.service.DeleteNote(c); err != nil {
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"status":  true,
		"message": "Deleted row successfully",
	})
}
