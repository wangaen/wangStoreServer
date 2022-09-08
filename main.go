package main

import (
	"fmt"
	"net/http"
)
import "github.com/gin-gonic/gin"

func main() {
	fmt.Println("项目初始化")
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    333,
		})
	})
	router.Run()
}
