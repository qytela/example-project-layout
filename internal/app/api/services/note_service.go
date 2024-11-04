package services

import (
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/qytela/example-project-layout/internal/app/api/models"
	"github.com/qytela/example-project-layout/internal/app/api/repository"
	"github.com/qytela/example-project-layout/internal/app/api/requests"
	"github.com/qytela/example-project-layout/internal/pkg/exception"
	"github.com/qytela/example-project-layout/internal/pkg/logger"
	"github.com/qytela/example-project-layout/internal/pkg/queryhelper"
	"github.com/qytela/example-project-layout/internal/pkg/utils"
)

type NoteService struct {
	repository *repository.NoteRepository
}

func NewNoteService(repository *repository.NoteRepository) *NoteService {
	return &NoteService{
		repository: repository,
	}
}

func (s *NoteService) GetUserNotes(c echo.Context) ([]models.UserNote, error) {
	userId := c.Get("userId").(uuid.UUID)

	data, err := s.repository.GetUserNotes(userId)
	if err != nil {
		logger.MakeLogEntry(nil).Info(err)
		return nil, exception.NewBadRequest()
	}

	return data, nil
}

func (s *NoteService) GetNotes(c echo.Context) ([]models.Note, error) {
	userId := c.Get("userId").(uuid.UUID)

	paramOptions := queryhelper.NewParamOptions()
	paramOptions.SetLimit(c.QueryParam("limit"))
	paramOptions.SetOffset(c.QueryParam("offset"))

	data, err := s.repository.GetNotes(userId, paramOptions)
	if err != nil {
		return nil, exception.NewBadRequest()
	}

	return data, nil
}

func (s *NoteService) StoreNote(c echo.Context) (models.Note, error) {
	userId := c.Get("userId").(uuid.UUID)

	req := new(requests.StoreNoteRequest)
	if err := utils.ValidateRequest(c, req); err != nil {
		return models.Note{}, exception.NewInvalidRequest(err)
	}

	data, err := s.repository.StoreNote(userId, req)
	if err != nil {
		logger.MakeLogEntry(c).Info(err)
		return models.Note{}, exception.NewBadRequest()
	}

	return data, nil
}

func (s *NoteService) UpdateNote(c echo.Context) (models.Note, error) {
	userId := c.Get("userId").(uuid.UUID)

	req := new(requests.UpdateNoteRequest)
	if err := utils.ValidateRequest(c, req); err != nil {
		return models.Note{}, exception.NewInvalidRequest(err)
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return models.Note{}, exception.NewBadRequest()
	}

	if _, err := s.repository.GetNote(userId, id); err != nil {
		return models.Note{}, exception.NewRecordNotFound()
	}

	data, err := s.repository.UpdateNote(userId, id, req)
	if err != nil {
		return models.Note{}, exception.NewBadRequest()
	}

	return data, nil
}

func (s *NoteService) DeleteNote(c echo.Context) error {
	userId := c.Get("userId").(uuid.UUID)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return exception.NewBadRequest()
	}

	if err := s.repository.DeleteNote(userId, id); err != nil {
		return exception.NewBadRequest()
	}

	return nil
}
