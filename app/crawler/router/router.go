package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wangStoreServer/app/crawler/engine"
	parses "wangStoreServer/app/crawler/parses/zhenaiwang/parseUserList"
)

func InitCrawlerRouter(e *gin.Engine) {
	e.GET("/crawler", func(c *gin.Context) {
		//eng := engine.ConcurrencyEngine{
		//	Scheduler:   &scheduler.SimpleScheduler{},
		//	WorkerCount: 100,
		//}
		//
		//eng.Run(engine.Request{
		//	Url:         "https://www.zhenai.com/zhenghun",
		//	ParseUrlFun: parses.ParseTag,
		//})

		engine.SimpleRun(engine.Request{
			Url:         "https://www.zhenai.com/zhenghun/cangzhou",
			ParseUrlFun: parses.ParseUserList,
		})

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
		})
	})
}
