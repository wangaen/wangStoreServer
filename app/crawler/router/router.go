package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wangStoreServer/app/crawler/fetcher"
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
		//file, err := os.OpenFile("/app/crawler/parses/parseCityList/cityList_test_data.html", os.O_RDWR|os.O_TRUNC, 0666)
		//if err != nil {
		//	panic(err)
		//	return
		//}
		//defer file.Close()
		bytes, _ := fetcher.Fetch("https://www.zhenai.com/zhenghun")
		//write := bufio.NewWriter(file)
		//write.Write(bytes)
		//write.Flush()
		//parseRequest := parses.ParseCityList(bytes)
		fmt.Println(string(bytes))

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
		})
	})
}
