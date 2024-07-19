package factory

import (
	"github.com/charmingruby/kickstart/internal/domain/example/example_entity"
	"github.com/charmingruby/kickstart/internal/domain/example/example_repository"
)

func MakeExample(
	repo example_repository.ExampleRepository,
	name string,
) (*example_entity.Example, error) {
	example, err := example_entity.NewExample(name)
	if err != nil {
		return nil, err
	}

	if err := repo.Store(example); err != nil {
		return nil, err
	}

	return example, nil
}
