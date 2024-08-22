package v1

import (
	_ "github.com/charmingruby/kickstart/api"
	"github.com/charmingruby/kickstart/internal/common/api/rest"
	"github.com/charmingruby/kickstart/internal/common/core/validation"
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
//	@Success		201		{object}	rest.Response
//	@Failure		400		{object}	rest.Response
//	@Failure		500		{object}	rest.Response
//	@Router			/examples [post]
func (h *Handler) createExampleEndpoint(c *gin.Context) {
	var req CreateExampleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		rest.NewPayloadError(c, err)
		return
	}

	dto := dto.CreateExampleUseCaseDTO{
		Name: req.Name,
	}

	if err := h.exampleService.CreateExampleUseCase(dto); err != nil {
		validationErr, ok := err.(*validation.ErrValidation)
		if ok {
			rest.NewEntityError(c, validationErr)
			return
		}

		rest.NewInternalServerError(c, err)
		return
	}
	rest.NewCreatedResponse(c, "example")
}
