package example

import (
	"github.com/charmingruby/kickstart/internal/example/domain/repository"
	"github.com/charmingruby/kickstart/internal/example/domain/usecase"
	v1 "github.com/charmingruby/kickstart/internal/example/transport/rest/endpoint/v1"
	"github.com/gin-gonic/gin"
)

func NewService(exampleRepo repository.ExampleRepository) *usecase.ExampleUseCaseRegistry {
	return usecase.NewExampleUseCaseRegistry(exampleRepo)
}

func NewHTTPService(router *gin.Engine, exampleService usecase.ExampleUseCase) *v1.Handler {
	return v1.NewHandler(router, exampleService)
}
