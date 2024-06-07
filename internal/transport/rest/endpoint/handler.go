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
	api := h.router.Group("/api")
	{
		api.GET("/welcome", welcomeEndpoint)

		api.POST("/examples", h.createExampleEndpoint)
		api.GET("/examples/:id", h.getExampleEndpoint)
	}
}
