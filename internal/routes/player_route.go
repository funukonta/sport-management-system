package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/sport-management-system/internal/handlers"
	"github.com/sport-management-system/internal/middleware"
	"github.com/sport-management-system/internal/repository"
	"github.com/sport-management-system/internal/services"
)

func PlayerRoutes(r *gin.RouterGroup, db *sqlx.DB) {
	playerRepo := repository.NewPlayerRepository(db)
	teamRepo := repository.NewTeamRepository(db)
	playerService := services.NewPlayerService(playerRepo, teamRepo)
	playerHandler := handlers.NewPlayerHandler(playerService)

	player := r.Group("/players")
	player.Use(middleware.AuthMiddleware())

	player.POST("/", playerHandler.Create)
	player.GET("/", playerHandler.FindAll)
	player.GET("/:id", playerHandler.FindByID)
	player.PUT("/:id", playerHandler.Update)
	player.DELETE("/:id", playerHandler.Delete)

}
