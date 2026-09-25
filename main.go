package main

import (
	serve "hospital-a/api/server"
	"hospital-a/configs"
	"hospital-a/database"
	"hospital-a/handler"
	"hospital-a/repository"
	"hospital-a/routes"
	"hospital-a/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	config := configs.NewConfig()

	client, err := database.NewSQLClient(database.Config{
		DBDriver:          config.Database.DatabaseDriver,
		DBSource:          config.Database.DatabaseSource,
		MaxOpenConns:      25,
		MaxIdleConns:      25,
		ConnMaxIdleTime:   15 * time.Minute,
		ConnectionTimeout: 5 * time.Second,
	})

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize database client")
		return
	}

	defer func() {
		if err := client.Close(); err != nil {
			log.Error().Msgf("Failed to close database client: %v", err)
		}
	}()

	staffRepo := repository.NewStaffRepository(client.DB)
	patientRepo := repository.NewPatientRepository(client.DB)

	staffService := service.NewStaffService(staffRepo)
	patientService := service.NewPatientService(patientRepo)

	staffHandler := handler.NewStaffHandler(staffService)
	patientHandler := handler.NewPatientHandler(patientService)

	cors := config.CorsNew()

	router := gin.Default()

	router.Use(cors)
	api := router.Group("/api")
	v1 := api.Group("/v1")

	routes.StaffRoute(v1, staffHandler)
	routes.PatientRoute(v1, patientHandler)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server := serve.NewServer(log.Logger, router, config)
	server.Serve()
}
