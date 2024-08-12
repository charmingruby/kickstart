package presenter

import (
	"time"

	"github.com/charmingruby/kickstart/internal/example/domain/entity"
)

func NewExamplePresenter(example entity.Example) ExamplePresenter {
	return ExamplePresenter{
		ID:        example.GetID(),
		Name:      example.GetName(),
		CreatedAt: example.GetCreatedAt(),
	}
}

type ExamplePresenter struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
