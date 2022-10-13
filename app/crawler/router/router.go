package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wangStoreServer/app/crawler/engine"
	"wangStoreServer/app/crawler/parses/zhenaiwang/parseCityList"
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
			Url:         "https://www.zhenai.com/zhenghun",
			ParseUrlFun: parses.ParseCityList,
		})

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
		})
	})
}
