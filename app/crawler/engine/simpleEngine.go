package engine

import (
	"fmt"
	"log"
	"wangStoreServer/app/crawler/fetcher"
)

// SimpleRun 单机版爬虫
func SimpleRun(seeds ...Request) {
	var requests []Request // 储存每一个任务，深度处理每一个页面下的 a 标签，直到没有

	//	初始化
	for _, seed := range seeds {
		requests = append(requests, seed)
	}
	fmt.Println("fsdfsdf", requests)
	//	逐个发起请求
	for len(requests) > 0 {
		request := requests[0]
		// 截取掉
		requests = requests[1:]

		log.Printf(""+
			"\n\n***************************************\n"+
			"请求URL: %s\n***************************************\n\n",
			request.Url)

		bytes, err := fetcher.Fetch(request.Url)
		if err != nil {
			log.Printf("请求 %s 异常, err: %s \n", request.Url, err.Error())
		}

		parseResult := request.ParseUrlFun(bytes)

		// 收集
		requests = append(requests, parseResult.RequestArray...)
	}
}
