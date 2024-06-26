package example_repository

import "github.com/charmingruby/kickstart/internal/domain/example/example_entity"

type ExampleRepository interface {
	Store(e *example_entity.Example) error
	FindByID(id string) (*example_entity.Example, error)
}
