package usecase

import (
	"github.com/charmingruby/kickstart/internal/common/core/custom_err"
	"github.com/charmingruby/kickstart/internal/example/domain/entity"
)

func (s *ExampleUseCaseRegistry) GetExampleUseCase(id string) (*entity.Example, error) {
	example, err := s.exampleRepo.FindByID(id)
	if err != nil {
		return nil, custom_err.NewNotFoundErr("example")
	}

	return example, nil
}
