package endpoint

import (
	_ "github.com/charmingruby/kickstart/docs"
	"github.com/charmingruby/kickstart/internal/domain/example/dto"
	"github.com/charmingruby/kickstart/internal/validation"
	"github.com/gin-gonic/gin"
)

type CreateExampleRequest struct {
	Name string `json:"name" binding:"required"`
}

// CreateExample godoc
//
//	@Summary		Create example
//	@Description	Create a new example
//	@Tags			Examples
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateExampleRequest	true	"Add Example"
//	@Success		201		{object}	Response
//	@Failure		400		{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/examples [post]
func (h *Handler) CreateExampleEndpoint(c *gin.Context) {
	var req CreateExampleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		newPayloadError(c, err)
		return
	}

	dto := dto.CreateExampleDTO{
		Name: req.Name,
	}

	if err := h.exampleService.CreateExample(dto); err != nil {
		validationErr, ok := err.(*validation.ErrValidation)
		if ok {
			newBadRequestError(c, validationErr)
			return
		}

		newInternalServerError(c, err)
		return
	}
	newCreatedResponse(c, "example")
}
