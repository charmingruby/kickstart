package endpoint

import (
	"github.com/charmingruby/kickstart/internal/domain/example"
	"github.com/gin-gonic/gin"
)

func NewHandler(router *gin.Engine, exampleService example.ExampleServiceContract) *Handler {
	return &Handler{
		router:         router,
		exampleService: exampleService,
	}
}

type Handler struct {
	router         *gin.Engine
	exampleService example.ExampleServiceContract
}

func (h *Handler) Register() {
	apiV1 := h.router.Group("/api/v1")
	{
		apiV1.GET("/welcome", welcomeEndpoint)
		apiV1.POST("/examples", h.createExampleEndpoint)
		apiV1.GET("/examples/:id", h.getExampleEndpoint)
	}
}
