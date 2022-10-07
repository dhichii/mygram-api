package helper

import "github.com/gin-gonic/gin"

func CreateMessageResponse(c *gin.Context, httpCode int, message interface{}) {
	c.JSON(httpCode, map[string]interface{}{
		"message": message,
	})
}
