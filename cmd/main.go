package main

import (
	"github.com/dfryer1193/thehardway/config"
	"github.com/dfryer1193/thehardway/middleware"
	"github.com/dfryer1193/thehardway/routes"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

	// Init config
	config.LoadConfig()

	router := gin.Default()

	// Attach middleware
	router.Use(middleware.LoggingMiddleware())
	router.Use(middleware.AuthMiddleware())
	router.Use(gin.Recovery())

	routes.SetupRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to start server")
	}
}
