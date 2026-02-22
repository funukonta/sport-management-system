package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sport-management-system/config"
	"github.com/sport-management-system/internal/middleware"
	"github.com/sport-management-system/internal/routes"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	r := gin.Default()

	r.Use(middleware.ErrorHandler())

	// Call route
	routes.Setup(r, db)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}
