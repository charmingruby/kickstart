package example_endpoint_v1

import (
	_ "github.com/charmingruby/kickstart/docs"
	"github.com/charmingruby/kickstart/internal/common/api/api_rest"
	"github.com/charmingruby/kickstart/internal/common/core"
	"github.com/charmingruby/kickstart/internal/example/domain/dto"
	"github.com/gin-gonic/gin"
)

type CreateExampleRequest struct {
	Name string `json:"name" binding:"required"`
}

// CreateExample godoc
//
//	@Summary		Creates an example
//	@Description	Creates an example
//	@Tags			Examples
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateExampleRequest	true	"Create Example Payload"
//	@Success		201		{object}	api_rest.Response
//	@Failure		400		{object}	api_rest.Response
//	@Failure		500		{object}	api_rest.Response
//	@Router			/examples [post]
func (h *Handler) createExampleEndpoint(c *gin.Context) {
	var req CreateExampleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api_rest.NewPayloadError(c, err)
		return
	}

	dto := dto.CreateExampleUseCaseDTO{
		Name: req.Name,
	}

	if err := h.exampleService.CreateExampleUseCase(dto); err != nil {
		validationErr, ok := err.(*core.ErrValidation)
		if ok {
			api_rest.NewEntityError(c, validationErr)
			return
		}

		api_rest.NewInternalServerError(c, err)
		return
	}
	api_rest.NewCreatedResponse(c, "example")
}
