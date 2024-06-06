package endpoint

import "github.com/gin-gonic/gin"

func NewHandler(router *gin.Engine) *Handler {
	return &Handler{
		router: router,
	}
}

type Handler struct {
	router *gin.Engine
}

func (h *Handler) Register() {
	api := h.router.Group("/api")
	{
		api.GET("/welcome", welcomeEndpoint)
	}
}
