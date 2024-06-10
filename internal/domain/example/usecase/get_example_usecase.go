package usecase

import (
	"github.com/charmingruby/kickstart/internal/domain/example/entity"
	"github.com/charmingruby/kickstart/internal/validation"
)

func (s *ExampleService) GetExample(id string) (*entity.Example, error) {
	example, err := s.exampleRepo.FindByID(id)
	if err != nil {
		return nil, validation.NewNotFoundErr("example")
	}

	return example, nil
}
