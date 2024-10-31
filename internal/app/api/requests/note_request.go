package requests

type StoreNoteRequest struct {
	Note  string `json:"note" validate:"required"`
	Order int    `json:"order" validate:"required"`
}

type UpdateNoteRequest struct {
	Note  string `json:"note" validate:"required"`
	Order int    `json:"order" validate:"required"`
}
