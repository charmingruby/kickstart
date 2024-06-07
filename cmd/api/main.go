package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/charmingruby/kickstart/internal/config"
	"github.com/charmingruby/kickstart/internal/database"
	"github.com/charmingruby/kickstart/internal/domain/example"
	"github.com/charmingruby/kickstart/internal/transport/rest"
	"github.com/charmingruby/kickstart/internal/transport/rest/endpoint"
	"github.com/charmingruby/kickstart/pkg/postgres"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Warn("CONFIGURATION: .env file not found")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("CONFIGURATION: %s", err.Error()))
		os.Exit(1)
	}

	db, err := postgres.NewPostgresConnection(cfg)
	if err != nil {
		slog.Error(fmt.Sprintf("DATABASE: %s", err.Error()))
		os.Exit(1)
	}

	router := gin.Default()

	initDependencies(db, router)

	server := rest.NewServer(router, cfg.ServerConfig.Port)
	if err := server.Start(); err != nil {
		slog.Error(fmt.Sprintf("REST SERVER: %s", err.Error()))
		os.Exit(1)
	}
}

func initDependencies(db *sqlx.DB, router *gin.Engine) {
	exampleRepo, err := database.NewPostgresExampleRepository(db)
	if err != nil {
		slog.Error(fmt.Sprintf("DATABASE REPOSITORY: %s", err.Error()))
		os.Exit(1)
	}

	exampleSvc := example.NewExampleService(exampleRepo)

	endpoint.NewHandler(router, exampleSvc).Register()
}
