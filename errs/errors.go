package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"` // if empty won't appear in json format
	Message string `json:"message"`
}

// func to return only the err message
// can call this function in the handlers/http port
func (e *AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewInternalServerError(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}
