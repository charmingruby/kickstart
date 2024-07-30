package usecase

import (
	"testing"

	"github.com/charmingruby/kickstart/internal/example/domain/entity"
	"github.com/charmingruby/kickstart/test/inmemory"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	exampleRepo *inmemory.InMemoryExampleRepository
	useCase     *ExampleUseCaseRegistry
}

// initial setup
func (s *Suite) SetupSuite() {
	s.exampleRepo = inmemory.NewInMemoryExampleRepository()
	s.useCase = NewExampleUseCaseRegistry(s.exampleRepo)
}

// executes before each test
func (s *Suite) SetupTest() {
	s.exampleRepo.Items = []entity.Example{}
}

// executes after each test
func (s *Suite) TearDownTest() {
	s.exampleRepo.Items = []entity.Example{}
}

// executes before each sub test
func (s *Suite) SetupSubTest() {
	s.exampleRepo.Items = []entity.Example{}
}

// executes after each sub test
func (s *Suite) TearDownSubTest() {
	s.exampleRepo.Items = []entity.Example{}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
