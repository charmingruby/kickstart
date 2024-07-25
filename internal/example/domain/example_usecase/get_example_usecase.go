package example_usecase

import (
	"github.com/charmingruby/kickstart/internal/common/core"
	"github.com/charmingruby/kickstart/internal/example/domain/example_entity"
)

func (s *ExampleUseCaseRegistry) GetExampleUseCase(id string) (*example_entity.Example, error) {
	example, err := s.exampleRepo.FindByID(id)
	if err != nil {
		return nil, core.NewNotFoundErr("example")
	}

	return example, nil
}
