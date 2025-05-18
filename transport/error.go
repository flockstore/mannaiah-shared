package transport

import (
	"errors"
	"fmt"
)

// AppError defines a generic, transport-agnostic error with status code and message.
type AppError struct {
	Code    int    // Code is the suggested HTTP status code
	Message string // Message is a Human-readable message
	Err     error  // Err is an Underlying error (optional)
}

// Error implements the error interface.
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// CodeOf extracts the status code from an error (default: 500).
func CodeOf(err error) int {
	var e *AppError
	if errors.As(err, &e) {
		return e.Code
	}
	return 500
}

// New creates a new error
func New(code int, msg string, err error) *AppError {
	return &AppError{Code: code, Message: msg, Err: err}
}

// BadRequest creates a bad request error
func BadRequest(msg string) *AppError {
	return New(400, msg, nil)
}

// NotFound creates a not found error
func NotFound(msg string) *AppError {
	return New(404, msg, nil)
}

// Conflict creates a conflict error
func Conflict(msg string) *AppError {
	return New(409, msg, nil)
}
