package endpoint

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func newResponse(c *gin.Context, code int, data any, message string) {
	res := Response{
		Message: message,
		Data:    data,
		Code:    code,
	}
	c.JSON(code, res)
}

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Code    int    `json:"status_code"`
}

func newCreatedResponse(c *gin.Context, entity string) {
	msg := entity + " created successfully"
	newResponse(c, http.StatusCreated, nil, msg)
}

func newOkResponse(c *gin.Context, msg string, data any) {
	newResponse(c, http.StatusOK, data, msg)
}

func newPayloadError(c *gin.Context, err error) {
	newResponse(c, http.StatusBadRequest, nil, "Payload error: "+err.Error())
}

func newBadRequestError(c *gin.Context, err error) {
	newResponse(c, http.StatusBadRequest, nil, err.Error())
}

func newResourceNotFoundError(c *gin.Context, err error) {
	newResponse(c, http.StatusNotFound, nil, err.Error())
}

func newInternalServerError(c *gin.Context, err error) {
	newResponse(c, http.StatusInternalServerError, nil, err.Error())
}
