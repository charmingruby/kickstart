package mapper

import (
	"time"

	"github.com/charmingruby/kickstart/internal/example/domain/entity"
)

func DomainExampleToPostgres(example entity.Example) PostgresExample {
	return PostgresExample{
		ID:        example.GetID(),
		Name:      example.GetName(),
		CreatedAt: example.GetCreatedAt(),
	}
}

func PostgresExampleToDomain(pgExample PostgresExample) entity.Example {
	example := entity.Example{}
	example.SetID(pgExample.ID)
	example.SetName(pgExample.Name)
	example.SetCreatedAt(pgExample.CreatedAt)
	return example
}

type PostgresExample struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
