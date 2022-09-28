package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wangStoreServer/app/crawler/engine"
	"wangStoreServer/app/crawler/parses"
	"wangStoreServer/app/crawler/scheduler"
)

func InitCrawlerRouter(e *gin.Engine) {
	e.GET("/crawler", func(c *gin.Context) {
		eng := engine.ConcurrencyEngine{
			Scheduler:   &scheduler.QueueScheduler{},
			WorkerCount: 100,
		}

		eng.Run(engine.Request{
			Url:         "https://book.douban.com/",
			ParseUrlFun: parses.ParseTag,
		})

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
		})
	})
}
