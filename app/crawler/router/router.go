package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wangStoreServer/app/crawler/engine"
	parses2 "wangStoreServer/app/crawler/parses/zhenaiwang/parseCityList"
	"wangStoreServer/app/crawler/persist"
	"wangStoreServer/app/crawler/scheduler"
)

func InitCrawlerRouter(e *gin.Engine) {
	e.GET("/crawler", func(c *gin.Context) {
		eng := engine.ConcurrencyEngine{
			Scheduler:   &scheduler.QueueScheduler{},
			WorkerCount: 100,
			ItemChan:    persist.ItemSave(),
		}

		eng.Run(engine.Request{
			Url:         "https://www.zhenai.com/zhenghun",
			ParseUrlFun: parses2.ParseCityList,
		})

		//engine.SimpleEngine{}.Run(engine.Request{
		//	Url:         "https://www.zhenai.com/zhenghun",
		//	ParseUrlFun: parses2.ParseCityList,
		//})

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
		})
	})
}
