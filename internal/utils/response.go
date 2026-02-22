package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Meta struct {
	Total   int    `json:"total,omitempty"`
	Page    int    `json:"page,omitempty"`
	PerPage int    `json:"per_page,omitempty"`
	Code    int    `json:"code,omitempty"`
	Status  string `json:"status,omitempty"`
}

type Response struct {
	Message string      `json:"message"`
	Meta    *Meta       `json:"meta,omitempty"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, Response{
		Message: message,
		Data:    data,
	})
}

func SuccessWithMeta(c *gin.Context, code int, message string, data interface{}, meta *Meta) {
	c.JSON(code, Response{
		Message: message,
		Meta:    meta,
		Data:    data,
	})
}

func Error(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Message: message,
		Meta: &Meta{
			Code:   code,
			Status: http.StatusText(code),
		},
		Data: nil,
	})
}

func ValidationError(c *gin.Context, errs interface{}) {
	c.JSON(http.StatusBadRequest, Response{
		Message: "validation error",
		Meta: &Meta{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
		},
		Data: errs,
	})
}
