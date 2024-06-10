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

//	@title			Kickstart
//	@version		1.0
//	@description	This is the kickstart for new Go APIs

//	@contact.name	Gustavo Dias
//	@contact.url	https://github.com/charmingruby
//	@contact.email	gustavodiasa2121@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:3000
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

//	@securityDefinitions.apikey	JWT
//	@in							header
//	@name						Authorization
//	@description				Description for what is this security definition being used

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
	exampleRepo, err := database.NewPostgresExampleRepository(db)
	if err != nil {
		slog.Error(fmt.Sprintf("DATABASE REPOSITORY: %s", err.Error()))
		os.Exit(1)
	}

	exampleSvc := example.NewExampleService(exampleRepo)

	endpoint.NewHandler(router, exampleSvc).Register()
}
