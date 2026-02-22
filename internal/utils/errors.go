package utils

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrNotFound            = errors.New("resource not found")
	ErrConflict            = errors.New("resource already exists")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("forbidden")
	ErrBadRequest          = errors.New("bad request")
	ErrInternalServerError = errors.New("internal server error")
)

type CustomError struct {
	Code    int
	Message string
	Data    interface{}
	Err     error
}

func (e *CustomError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	if e.Err != nil {
		return e.Err.Error()
	}
	return http.StatusText(e.Code)
}

func (e *CustomError) Unwrap() error {
	return e.Err
}

func NewHTTPError(code int, message string, err error) *CustomError {
	return &CustomError{Code: code, Message: message, Err: err}
}

func NewBadRequestError(message string) *CustomError {
	return NewHTTPError(http.StatusBadRequest, message, ErrBadRequest)
}

func NewUnauthorizedError(message string) *CustomError {
	return NewHTTPError(http.StatusUnauthorized, message, ErrUnauthorized)
}

func NewInternalServerError(message string) *CustomError {
	return NewHTTPError(http.StatusInternalServerError, message, ErrInternalServerError)
}

func NewValidationError(details interface{}) *CustomError {
	return &CustomError{
		Code:    http.StatusBadRequest,
		Message: "validation error",
		Data:    details,
		Err:     ErrBadRequest,
	}
}

func IsUniqueDataError(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" {
			return true
		}
	}
	return false
}
