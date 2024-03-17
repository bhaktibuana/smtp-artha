package helpers

import "github.com/gin-gonic/gin"

type S_Response struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}

func Response(message string, httpStatus int, context *gin.Context, data interface{}) {
	response := S_Response{
		Message: message,
		Status:  httpStatus >= 200 && httpStatus < 300,
		Data:    data,
	}

	if response.Status {
		context.JSON(httpStatus, response)
	} else {
		context.AbortWithStatusJSON(httpStatus, response)
	}
}
