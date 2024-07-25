package example

import (
	"github.com/charmingruby/kickstart/internal/example/domain/example_repository"
	"github.com/charmingruby/kickstart/internal/example/domain/example_usecase"
	"github.com/charmingruby/kickstart/internal/example/transport/rest/endpoint/example_endpoint_v1"
	"github.com/gin-gonic/gin"
)

func NewService(exampleRepo example_repository.ExampleRepository) *example_usecase.ExampleUseCaseRegistry {
	return example_usecase.NewExampleUseCaseRegistry(exampleRepo)
}

func NewHTTPService(router *gin.Engine, exampleService example_usecase.ExampleUseCase) *example_endpoint_v1.Handler {
	return example_endpoint_v1.NewHandler(router, exampleService)
}
