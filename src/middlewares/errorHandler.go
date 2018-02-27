package middlewares

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// ErrorHandler custom handle error
func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		errorDetail := strings.Split(c.Errors[0].Error(), "|")
		statusCode, _ := strconv.Atoi(errorDetail[0])
		c.JSON(statusCode, gin.H{
			"body": errorDetail[1],
		})
	}
}
