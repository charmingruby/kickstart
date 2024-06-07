package endpoint

import (
	"github.com/charmingruby/kickstart/internal/domain/example"
	"github.com/charmingruby/kickstart/internal/validation"
	"github.com/gin-gonic/gin"
)

type createExampleRequest struct {
	Name string `json:"name" binding:"required"`
}

func (h *Handler) createExampleEndpoint(c *gin.Context) {
	var req createExampleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		newPayloadError(c, err)
		return
	}

	dto := example.CreateExampleDTO{
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
