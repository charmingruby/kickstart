package integration

import (
	"fmt"
	"log/slog"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/charmingruby/kickstart/internal/common/api/api_rest"
	"github.com/charmingruby/kickstart/internal/example/database/postgres_repository"
	"github.com/charmingruby/kickstart/internal/example/domain/repository"
	"github.com/charmingruby/kickstart/internal/example/domain/usecase"
	v1 "github.com/charmingruby/kickstart/internal/example/transport/rest/endpoint/v1"
	"github.com/charmingruby/kickstart/test/container"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

const (
	contentType = "application/json"
)

type Suite struct {
	suite.Suite
	container   *container.TestDatabase
	server      *httptest.Server
	handler     *v1.Handler
	exampleRepo repository.ExampleRepository
}

func (s *Suite) SetupSuite() {
	tdb := container.NewPostgresTestDatabase()
	s.container = tdb
}

func (s *Suite) TearDownSuite() {
	s.container.DB.Close()
}

func (s *Suite) SetupSubTest() {
	err := s.container.RunMigrations()
	s.NoError(err)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	s.exampleRepo, err = postgres_repository.NewPostgresExampleRepository(s.container.DB)
	if err != nil {
		slog.Error(fmt.Sprintf("INTEGRATION TEST, DATABASE REPOSITORY: %s", err.Error()))
		os.Exit(1)
	}

	exampleSvc := usecase.NewExampleUseCaseRegistry(s.exampleRepo)
	s.handler = v1.NewHandler(router, exampleSvc)
	s.handler.Register()
	server := api_rest.NewServer(router, "3000")

	s.server = httptest.NewServer(server.Router)
}

func (s *Suite) TearDownSubTest() {
	err := s.container.RollbackMigrations()
	s.NoError(err)

	s.server.Close()
}

func (s *Suite) V1Route(path string) string {
	return fmt.Sprintf("%s/api/v1%s", s.server.URL, path)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
