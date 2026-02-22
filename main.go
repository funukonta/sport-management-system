package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sport-management-system/config"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	r := gin.Default()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	srv.ListenAndServe()
}
