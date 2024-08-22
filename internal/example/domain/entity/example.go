package entity

import (
	"time"

	"github.com/charmingruby/kickstart/internal/common/core"
	"github.com/charmingruby/kickstart/internal/common/core/validation"
)

func NewExample(name string) (*Example, error) {
	p := ExamplePayload{
		ID:        core.NewID(),
		Name:      name,
		CreatedAt: time.Now(),
	}

	if err := validation.ValidateStruct(p); err != nil {
		return nil, err
	}

	example := Example{
		id:        p.ID,
		name:      p.Name,
		createdAt: p.CreatedAt,
	}

	return &example, nil
}

type ExamplePayload struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=3,max=16"`
	CreatedAt time.Time `validate:"required"`
}

type Example struct {
	id        string    `validate:"required"`
	name      string    `validate:"min=3,max=16"`
	createdAt time.Time `validate:"required"`
}

func (e *Example) GetID() string {
	return e.id
}

func (e *Example) SetID(id string) {
	e.id = id
}

func (e *Example) GetName() string {
	return e.name
}

func (e *Example) SetName(name string) {
	e.name = name
}

func (e *Example) GetCreatedAt() time.Time {
	return e.createdAt
}

func (e *Example) SetCreatedAt(createdAt time.Time) {
	e.createdAt = createdAt
}
