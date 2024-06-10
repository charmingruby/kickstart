package integration

import (
	"fmt"
	"log/slog"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/charmingruby/kickstart/internal/database"
	"github.com/charmingruby/kickstart/internal/domain/example/repository"
	"github.com/charmingruby/kickstart/internal/domain/example/usecase"
	"github.com/charmingruby/kickstart/internal/transport/rest"
	"github.com/charmingruby/kickstart/internal/transport/rest/endpoint"
	"github.com/charmingruby/kickstart/test/container"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	contentType = "application/json"
)

type Suite struct {
	suite.Suite
	container   *container.TestDatabase
	server      *httptest.Server
	handler     *endpoint.Handler
	exampleRepo repository.ExampleRepository
}

func (s *Suite) SetupSuite() {
	tdb := container.NewPostgresTestDatabase()
	s.container = tdb
}

func (s *Suite) TearDownSuite() {
	s.container.DB.Close()
}

func (s *Suite) SetupTest() {
	err := s.container.RunMigrations()
	assert.NoError(s.T(), err)

	router := gin.Default()
	s.exampleRepo, err = database.NewPostgresExampleRepository(s.container.DB)
	if err != nil {
		slog.Error(fmt.Sprintf("INTEGRATION TEST, DATABASE REPOSITORY: %s", err.Error()))
		os.Exit(1)
	}

	exampleSvc := usecase.NewExampleService(s.exampleRepo)
	s.handler = endpoint.NewHandler(router, exampleSvc)
	s.handler.Register()
	server := rest.NewServer(router, "3000")

	s.server = httptest.NewServer(server.Router)
}

func (s *Suite) TearDownTest() {
	err := s.container.RollbackMigrations()
	assert.NoError(s.T(), err)

	s.server.Close()
}

func (s *Suite) Route(path string) string {
	return fmt.Sprintf("%s/api%s", s.server.URL, path)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
