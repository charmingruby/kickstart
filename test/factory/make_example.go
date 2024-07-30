package factory

import (
	"github.com/charmingruby/kickstart/internal/example/domain/entity"
	"github.com/charmingruby/kickstart/internal/example/domain/repository"
)

func MakeExample(
	repo repository.ExampleRepository,
	name string,
) (*entity.Example, error) {
	example, err := entity.NewExample(name)
	if err != nil {
		return nil, err
	}

	if err := repo.Store(example); err != nil {
		return nil, err
	}

	return example, nil
}
