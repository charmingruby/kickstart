package v1

import (
	"github.com/charmingruby/kickstart/internal/common/api/rest"
	"github.com/charmingruby/kickstart/internal/common/core/custom_err"
	"github.com/charmingruby/kickstart/internal/example/transport/rest/presenter"
	"github.com/gin-gonic/gin"
)

type GetExampleResponse struct {
	Message string                      `json:"message"`
	Data    *presenter.ExamplePresenter `json:"data"`
	Code    int                         `json:"status_code"`
}

// GetExample godoc
//
//	@Summary		Gets an example
//	@Description	Gets an example
//	@Tags			Examples
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Get Example Payload"
//	@Success		200	{object}	GetExampleResponse
//	@Failure		404	{object}	rest.Response
//	@Failure		500	{object}	rest.Response
//	@Router			/examples/{id} [get]
func (h *Handler) getExampleEndpoint(c *gin.Context) {
	exampleID := c.Param("id")

	example, err := h.exampleService.GetExampleUseCase(exampleID)
	if err != nil {
		resourceNotFoundErr, ok := err.(*custom_err.ErrNotFound)
		if ok {
			rest.NewResourceNotFoundError(c, resourceNotFoundErr)
			return
		}

		rest.NewInternalServerError(c, err)
		return
	}

	examplePresenter := presenter.NewExamplePresenter(*example)

	rest.NewOkResponse(
		c,
		"example found",
		examplePresenter,
	)
}
