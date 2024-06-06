package example

import (
	"log/slog"

	"github.com/charmingruby/kickstart/internal/validation"
)

type ExampleServiceContract interface {
	CreateExample(dto CreateExampleDTO) error
	GetExample(id string) (*Example, error)
}

func NewExampleService(exampleRepo ExampleRepository) *ExampleService {
	return &ExampleService{
		exampleRepo: exampleRepo,
	}
}

type ExampleService struct {
	exampleRepo ExampleRepository
}

func (s *ExampleService) CreateExample(dto CreateExampleDTO) error {
	example, err := NewExample(dto.Name)
	if err != nil {
		return err
	}

	if err := s.exampleRepo.Store(example); err != nil {
		slog.Error(err.Error())
		return validation.NewInternalErr()
	}

	return nil
}

func (s *ExampleService) GetExample(id string) (*Example, error) {
	example, err := s.exampleRepo.FindByID(id)
	if err != nil {
		return nil, validation.NewNotFoundErr("example")
	}

	return example, nil
}
