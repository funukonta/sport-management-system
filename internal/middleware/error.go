package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sport-management-system/internal/utils"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 || c.Writer.Written() {
			return
		}

		err := c.Errors.Last().Err
		code, message, data := MapError(err)

		c.JSON(code, utils.Response{
			Message: message,
			Meta: &utils.Meta{
				Code:   code,
				Status: http.StatusText(code),
			},
			Data: data,
		})
	}
}

func MapError(err error) (int, string, interface{}) {
	var httpErr *utils.HTTPError
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
	case errors.Is(err, utils.ErrNotFound):
		return http.StatusNotFound, "data not found", nil
	case errors.Is(err, utils.ErrConflict):
		return http.StatusConflict, "data already exists", nil
	case errors.Is(err, utils.ErrUnauthorized):
		return http.StatusUnauthorized, "unauthorized", nil
	case errors.Is(err, utils.ErrForbidden):
		return http.StatusForbidden, "forbidden", nil
	case errors.Is(err, utils.ErrBadRequest):
		return http.StatusBadRequest, err.Error(), nil
	default:
		return http.StatusInternalServerError, "internal server error", nil
	}
}
