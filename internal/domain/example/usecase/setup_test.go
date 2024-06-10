package usecase

import (
	"testing"

	"github.com/charmingruby/kickstart/internal/domain/example/entity"
	"github.com/charmingruby/kickstart/test/inmemory"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	exampleRepo    *inmemory.InMemoryExampleRepository
	exampleService *ExampleService
}

func (s *Suite) SetupSuite() {
	s.exampleRepo = inmemory.NewInMemoryExampleRepository()
	s.exampleService = NewExampleService(s.exampleRepo)
}

func (s *Suite) SetupTest() {
	s.exampleRepo.Items = []entity.Example{}
}

func (s *Suite) TearDownTest() {
	s.exampleRepo.Items = []entity.Example{}
}

func (s *Suite) SetupSubTest() {
	s.exampleRepo.Items = []entity.Example{}
}

func (s *Suite) TearDownSubTest() {
	s.exampleRepo.Items = []entity.Example{}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
