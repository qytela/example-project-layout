package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/qytela/example-project-layout/internal/app/api/models"
	"github.com/qytela/example-project-layout/internal/app/api/requests"
	"github.com/qytela/example-project-layout/internal/pkg/queryhelper"
)

type NoteRepository struct {
	db *sqlx.DB
}

func NewNoteRepository(db *sqlx.DB) *NoteRepository {
	return &NoteRepository{
		db: db,
	}
}

func (r *NoteRepository) GetNotes(userId uuid.UUID, paramOptions *queryhelper.ParamOptions) ([]models.Note, error) {
	var notes []models.Note

	query := `SELECT * FROM notes WHERE user_id = $1 LIMIT $2 OFFSET $3`
	if err := r.db.Select(&notes, query, userId, paramOptions.Limit, paramOptions.Offset); err != nil {
		return notes, err
	}

	return notes, nil
}

func (r *NoteRepository) GetNote(userId uuid.UUID, id int) (models.Note, error) {
	var note models.Note

	query := `SELECT * FROM notes WHERE user_id = $1 AND id = $2`
	if err := r.db.Get(&note, query, userId, id); err != nil {
		return note, err
	}

	return note, nil
}

func (r *NoteRepository) StoreNote(userId uuid.UUID, req *requests.StoreNoteRequest) (models.Note, error) {
	var note models.Note

	query := `INSERT INTO notes (user_id, note, "order") VALUES ($1, $2, $3) RETURNING *`
	if err := r.db.
		QueryRowx(query, userId, req.Note, req.Order).
		StructScan(&note); err != nil {
		return note, err
	}

	return note, nil
}

func (r *NoteRepository) UpdateNote(userId uuid.UUID, id int, req *requests.UpdateNoteRequest) (models.Note, error) {
	var note models.Note

	query := `
		UPDATE notes
		SET note = $1, "order" = $2 
		WHERE user_id = $3 AND id = $4
		RETURNING *
	`

	qry, args, err := sqlx.In(query, req.Note, req.Order, userId, id)
	if err != nil {
		return note, err
	}

	if err := r.db.Get(&note, qry, args...); err != nil {
		return note, err
	}

	return note, nil
}

func (r *NoteRepository) DeleteNote(userId uuid.UUID, id int) error {
	query := `DELETE FROM notes WHERE user_id = $1 AND id = $2`
	if _, err := r.db.Exec(query, userId, id); err != nil {
		return err
	}

	return nil
}
