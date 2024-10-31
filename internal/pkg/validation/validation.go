package validation

import (
	"github.com/go-playground/validator/v10"
)

type Validation struct {
	validation *validator.Validate
}

func NewValidation() *Validation {
	validation := validator.New()
	return &Validation{validation}
}

func (v *Validation) Validate(i interface{}) error {
	return v.validation.Struct(i)
}
