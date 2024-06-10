package usecase

import (
	"github.com/charmingruby/kickstart/internal/domain/example/dto"
	"github.com/charmingruby/kickstart/internal/domain/example/entity"
	"github.com/charmingruby/kickstart/internal/domain/example/repository"
)

type ExampleServiceContract interface {
	CreateExample(dto dto.CreateExampleDTO) error
	GetExample(id string) (*entity.Example, error)
}

func NewExampleService(exampleRepo repository.ExampleRepository) *ExampleService {
	return &ExampleService{
		exampleRepo: exampleRepo,
	}
}

type ExampleService struct {
	exampleRepo repository.ExampleRepository
}
