package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/sport-management-system/internal/handlers"
	"github.com/sport-management-system/internal/repository"
	"github.com/sport-management-system/internal/services"
)

func AuthRoutes(r *gin.RouterGroup, db *sqlx.DB) {
	authRepo := repository.NewUserRepository(db)
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	r.POST("/login", authHandler.Login)
}
