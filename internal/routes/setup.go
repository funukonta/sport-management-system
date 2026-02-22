package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Setup(r *gin.Engine, db *sqlx.DB) {
	api := r.Group("/api")

	v1 := api.Group("/v1")

	AuthRoutes(v1, db)
	TeamRoutes(v1, db)
}
