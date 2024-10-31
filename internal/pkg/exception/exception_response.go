package exception

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func NewInvalidRequest(err interface{}) *ErrorResponse {
	type Error struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return NewRequestMalformed()
	}

	var errors []Error
	for _, err := range validationErrors {
		field := err.Field()
		message := fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", field, err.Tag())
		errors = append(errors, Error{
			Field:   field,
			Message: message,
		})
	}

	return &ErrorResponse{
		Status:         false,
		Code:           400,
		Message:        "Invalid Request",
		InvalidRequest: errors,
	}
}

func NewRecordNotFound() *ErrorResponse {
	return &ErrorResponse{
		Status:  false,
		Code:    404,
		Message: "Record not found",
	}
}

func NewBadRequest() *ErrorResponse {
	return &ErrorResponse{
		Status:  false,
		Code:    400,
		Message: "Bad Request",
	}
}

func NewRequestMalformed() *ErrorResponse {
	return &ErrorResponse{
		Status:  false,
		Code:    400,
		Message: "Request Malformed",
	}
}

func NewUnauthorized() *ErrorResponse {
	return &ErrorResponse{
		Status:  false,
		Code:    401,
		Message: "Unauthorized",
	}
}

func NewNotVerifiedUser() *ErrorResponse {
	return &ErrorResponse{
		Status:  false,
		Code:    401,
		Message: "Your account has not been verified",
	}
}

func NewHandlePanic(err error) *ErrorResponse {
	res := &ErrorResponse{
		Status:     false,
		Code:       500,
		Message:    "ERROR",
		Exceptions: err,
	}

	return res
}
