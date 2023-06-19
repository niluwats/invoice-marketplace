package errors

import "net/http"

type AppError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

// AsMessage returns error message as a pointer of AppError
func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

// NewNotFoundError returns 404 http.StatusNotFound error
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

// NewUnexpectedError returns 500 http.StatusInternalServerError error
func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

// NewValidationError returns 422 http.StatusUnprocessableEntity error
func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

// NewBadRequest returns 422 http.StatusBadRequest error
func NewBadRequest(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

// NewAuthenticationError returns 401 http.StatusUnauthorized error
func NewAuthenticationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

// NewForbiddenError returns 403 http.StatusForbidden error
func NewForbiddenError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusForbidden,
	}
}

// NewConflictError returns 403 http.StatusConflict error
func NewConflictError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusConflict,
	}
}
