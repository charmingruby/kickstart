package example_usecase

import (
	"github.com/charmingruby/kickstart/internal/core"
	"github.com/charmingruby/kickstart/internal/domain/example/example_dto"
	"github.com/charmingruby/kickstart/internal/domain/example/example_entity"
)

func (s *ExampleService) CreateExampleUseCase(dto example_dto.CreateExampleUseCaseDTO) error {
	example, err := example_entity.NewExample(dto.Name)
	if err != nil {
		return err
	}

	if err := s.exampleRepo.Store(example); err != nil {
		return core.NewInternalErr("create example store")
	}

	return nil
}
