package inmemory

import (
	"github.com/charmingruby/kickstart/internal/common/core/custom_err"
	"github.com/charmingruby/kickstart/internal/example/domain/entity"
)

func NewInMemoryExampleRepository() *InMemoryExampleRepository {
	return &InMemoryExampleRepository{
		Items: []entity.Example{},
	}
}

type InMemoryExampleRepository struct {
	Items []entity.Example
}

func (r *InMemoryExampleRepository) Store(e *entity.Example) error {
	r.Items = append(r.Items, *e)
	return nil
}

func (r *InMemoryExampleRepository) FindByID(id string) (*entity.Example, error) {
	for _, e := range r.Items {
		if e.GetID() == id {
			return &e, nil
		}
	}

	return nil, custom_err.NewNotFoundErr("example")
}
