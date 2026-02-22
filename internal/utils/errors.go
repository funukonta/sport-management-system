package utils

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound     = errors.New("resource not found")
	ErrConflict     = errors.New("resource already exists")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
	ErrBadRequest   = errors.New("bad request")
)

type HTTPError struct {
	Code    int
	Message string
	Data    interface{}
	Err     error
}

func (e *HTTPError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	if e.Err != nil {
		return e.Err.Error()
	}
	return http.StatusText(e.Code)
}

func (e *HTTPError) Unwrap() error {
	return e.Err
}

func NewHTTPError(code int, message string, err error) *HTTPError {
	return &HTTPError{Code: code, Message: message, Err: err}
}

func NewBadRequestError(message string) *HTTPError {
	return NewHTTPError(http.StatusBadRequest, message, ErrBadRequest)
}

func NewUnauthorizedError(message string) *HTTPError {
	return NewHTTPError(http.StatusUnauthorized, message, ErrUnauthorized)
}

func NewValidationError(details interface{}) *HTTPError {
	return &HTTPError{
		Code:    http.StatusBadRequest,
		Message: "validation error",
		Data:    details,
		Err:     ErrBadRequest,
	}
}

func MapError(err error) (int, string, interface{}) {
	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		code := httpErr.Code
		if code == 0 {
			code = http.StatusInternalServerError
		}

		message := httpErr.Message
		if message == "" {
			message = http.StatusText(code)
		}

		return code, message, httpErr.Data
	}

	switch {
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound, "resource not found", nil
	case errors.Is(err, ErrConflict):
		return http.StatusConflict, "resource already exists", nil
	case errors.Is(err, ErrUnauthorized):
		return http.StatusUnauthorized, "unauthorized", nil
	case errors.Is(err, ErrForbidden):
		return http.StatusForbidden, "forbidden", nil
	case errors.Is(err, ErrBadRequest):
		return http.StatusBadRequest, err.Error(), nil
	default:
		return http.StatusInternalServerError, "internal server error", nil
	}
}
