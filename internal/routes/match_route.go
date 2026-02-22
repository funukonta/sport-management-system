package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/sport-management-system/internal/handlers"
	"github.com/sport-management-system/internal/middleware"
	"github.com/sport-management-system/internal/repository"
	"github.com/sport-management-system/internal/services"
)

func MatchRoutes(r *gin.RouterGroup, db *sqlx.DB) {
	matchRepo := repository.NewMatchRepository(db)
	teamRepo := repository.NewTeamRepository(db)
	playerRepo := repository.NewPlayerRepository(db)
	matchLogRepo := repository.NewMatchLogRepository(db)
	matchService := services.NewMatchService(matchRepo, teamRepo, playerRepo, matchLogRepo)
	matchHandler := handlers.NewMatchHandler(matchService)

	match := r.Group("/matches")
	match.Use(middleware.AuthMiddleware())

	match.POST("/", matchHandler.Create)
	match.GET("/", matchHandler.FindAll)
	match.GET("/:id", matchHandler.FindByID)
	match.PUT("/:id", matchHandler.Update)
	match.DELETE("/:id", matchHandler.Delete)
	match.POST("/:id/logs", matchHandler.AddMatchLog)
	match.POST("/:id/finish", matchHandler.FinishMatch)
	match.GET("/:id/report", matchHandler.GetMatchReport)
}
