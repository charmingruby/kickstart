package endpoint

import (
	_ "github.com/charmingruby/kickstart/docs"
	"github.com/charmingruby/kickstart/internal/core"
	"github.com/charmingruby/kickstart/internal/domain/example/example_dto"
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
//	@Success		201		{object}	Response
//	@Failure		400		{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/examples [post]
func (h *Handler) createExampleEndpoint(c *gin.Context) {
	var req CreateExampleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		NewPayloadError(c, err)
		return
	}

	dto := example_dto.CreateExampleDTO{
		Name: req.Name,
	}

	if err := h.exampleService.CreateExample(dto); err != nil {
		validationErr, ok := err.(*core.ErrValidation)
		if ok {
			NewEntityError(c, validationErr)
			return
		}

		NewInternalServerError(c, err)
		return
	}
	NewCreatedResponse(c, "example")
}
