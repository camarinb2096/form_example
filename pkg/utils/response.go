package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSuccess(c *gin.Context, message string, data interface{}) {
	responseData := response{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	c.JSON(http.StatusOK, responseData)
}

func HandleError(c *gin.Context, status int, err string) {
	responseData := response{
		Status:  "error",
		Message: err,
		Data:    nil,
	}
	c.JSON(http.StatusBadRequest, responseData)
}
