package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"wangStoreServer/app/crawler/fetcher"
)

func InitCrawlerRouter(e *gin.Engine) {
	e.GET("/crawler", func(c *gin.Context) {
		var url = "https://news.163.com/"
		byteArr, err := fetcher.Fetch(url)
		if err != nil {
			log.Println("Fetch Error,", err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  string(byteArr),
		})
	})
}
