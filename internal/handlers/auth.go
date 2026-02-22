package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sport-management-system/internal/dto"
	"github.com/sport-management-system/internal/services"
	"github.com/sport-management-system/internal/utils"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(utils.NewValidationError(err.Error()))
		return
	}

	token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		c.Error(err)
		return
	}

	utils.Success(c, http.StatusOK, "login successful", dto.LoginResponse{
		Token: token,
	})
}
