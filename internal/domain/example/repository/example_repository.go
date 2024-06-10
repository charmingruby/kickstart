package repository

import "github.com/charmingruby/kickstart/internal/domain/example/entity"

type ExampleRepository interface {
	Store(e *entity.Example) error
	FindByID(id string) (*entity.Example, error)
}
