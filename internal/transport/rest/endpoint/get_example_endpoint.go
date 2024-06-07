package endpoint

import (
	"github.com/charmingruby/kickstart/internal/validation"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getExampleEndpoint(c *gin.Context) {
	exampleID := c.Param("id")

	example, err := h.exampleService.GetExample(exampleID)
	if err != nil {
		resourceNotFoundErr, ok := err.(*validation.ErrNotFound)
		if ok {
			newResourceNotFoundError(c, resourceNotFoundErr)
			return
		}

		newInternalServerError(c, err)
		return
	}

	newOkResponse(
		c,
		"example found",
		example,
	)
}
