package errors

import (
	"fmt"
	"net/http"
)

// ErrorType represents HTTP status or custom error code
type ErrorType uint

const (
	NoType         ErrorType = http.StatusInternalServerError
	InternalServer ErrorType = http.StatusInternalServerError
	BadRequest     ErrorType = http.StatusBadRequest
	Unauthorized   ErrorType = http.StatusUnauthorized
	NotFound       ErrorType = http.StatusNotFound
	Validation     ErrorType = http.StatusUnprocessableEntity
	ExtendedError  ErrorType = 50001
)

// AppError defines a structured error
type AppError struct {
	Type   ErrorType
	Err    error
	Fields map[string][]string // optional field validation errors
}

func (e AppError) Error() string {
	return e.Err.Error()
}

// Factory methods
func New(t ErrorType, msg string) error {
	return AppError{Type: t, Err: fmt.Errorf(msg)}
}

func AddFieldError(err error, field, message string) error {
	appErr, ok := err.(AppError)
	if !ok {
		appErr = AppError{Type: NoType, Err: err, Fields: make(map[string][]string)}
	}
	if appErr.Fields == nil {
		appErr.Fields = make(map[string][]string)
	}
	appErr.Fields[field] = append(appErr.Fields[field], message)
	return appErr
}

func GetFields(err error) map[string][]string {
	if appErr, ok := err.(AppError); ok {
		return appErr.Fields
	}
	return nil
}

func GetType(err error) ErrorType {
	if appErr, ok := err.(AppError); ok {
		return appErr.Type
	}
	return NoType
}
