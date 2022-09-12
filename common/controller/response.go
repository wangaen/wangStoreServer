package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrResponseBody(c *gin.Context, code int, status string, msg string) {
	c.JSON(code, gin.H{
		"status": status,
		"msg":    msg,
	})
}

func OkResponseBody(c *gin.Context, status string, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"msg":    msg,
		"data":   data,
	})
}
