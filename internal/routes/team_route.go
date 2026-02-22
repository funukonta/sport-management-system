package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/sport-management-system/internal/handlers"
	"github.com/sport-management-system/internal/middleware"
	"github.com/sport-management-system/internal/repository"
	"github.com/sport-management-system/internal/services"
)

func TeamRoutes(r *gin.RouterGroup, db *sqlx.DB) {
	teamRepo := repository.NewTeamRepository(db)
	teamService := services.NewTeamService(teamRepo)
	teamHandler := handlers.NewTeamHandler(teamService)

	playerRepo := repository.NewPlayerRepository(db)
	playerService := services.NewPlayerService(playerRepo, teamRepo)
	playerHandler := handlers.NewPlayerHandler(playerService)

	team := r.Group("/teams")
	team.Use(middleware.AuthMiddleware())

	team.POST("/", teamHandler.Create)
	team.GET("/", teamHandler.FindAll)
	team.GET("/:id", teamHandler.FindByID)
	team.PUT("/:id", teamHandler.Update)
	team.DELETE("/:id", teamHandler.Delete)
	team.GET("/:id/players", playerHandler.FindByTeamID)

}
