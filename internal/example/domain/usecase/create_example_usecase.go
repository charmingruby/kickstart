package usecase

import (
	"github.com/charmingruby/kickstart/internal/common/core/custom_err"
	"github.com/charmingruby/kickstart/internal/example/domain/dto"
	"github.com/charmingruby/kickstart/internal/example/domain/entity"
)

func (s *ExampleUseCaseRegistry) CreateExampleUseCase(dto dto.CreateExampleUseCaseDTO) error {
	example, err := entity.NewExample(dto.Name)
	if err != nil {
		return err
	}

	if err := s.exampleRepo.Store(example); err != nil {
		return custom_err.NewInternalErr("create example store")
	}

	return nil
}
