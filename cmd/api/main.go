package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmingruby/kickstart/config"
	"github.com/charmingruby/kickstart/internal/common/api/api_rest"
	"github.com/charmingruby/kickstart/internal/example"
	"github.com/charmingruby/kickstart/internal/example/database/postgres_repository"
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
	api_rest.SetupCORS(router)

	initDependencies(db, router)

	server := api_rest.NewServer(router, cfg.ServerConfig.Port)

	go func() {
		if err := server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error(fmt.Sprintf("REST SERVER: %s", err.Error()))
			os.Exit(1)
		}
	}()

	// Graceful shutdown
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	slog.Info("HTTP Server interruption received!")

	ctx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Server.Shutdown(ctx); err != nil {
		slog.Error(fmt.Sprintf("GRACEFUL SHUTDOWN REST SERVER: %s", err.Error()))
		os.Exit(1)
	}

	slog.Info("Gracefully shutdown!")
}

func initDependencies(db *sqlx.DB, router *gin.Engine) {
	// Example module
	exampleRepo, err := postgres_repository.NewPostgresExampleRepository(db)
	if err != nil {
		slog.Error(fmt.Sprintf("DATABASE REPOSITORY: %s", err.Error()))
		os.Exit(1)
	}
	exampleSvc := example.NewService(exampleRepo)
	example.NewHTTPService(router, exampleSvc).Register()
}
