package endpoint

import (
	"github.com/charmingruby/kickstart/internal/core"
	"github.com/gin-gonic/gin"
)

// GetExample godoc
//
//	@Summary		Get example
//	@Description	Find an example
//	@Tags			Examples
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Example UUID"
//	@Success		200	{object}	Response
//	@Failure		404	{object}	Response
//	@Failure		500	{object}	Response
//	@Router			/examples/{id} [get]
func (h *Handler) getExampleEndpoint(c *gin.Context) {
	exampleID := c.Param("id")

	example, err := h.exampleService.GetExample(exampleID)
	if err != nil {
		resourceNotFoundErr, ok := err.(*core.ErrNotFound)
		if ok {
			NewResourceNotFoundError(c, resourceNotFoundErr)
			return
		}

		NewInternalServerError(c, err)
		return
	}

	NewOkResponse(
		c,
		"example found",
		example,
	)
}
