package engine

import (
	"fmt"
	"log"
	"wangStoreServer/app/crawler/fetcher"
)

type SimpleEngine struct {
}

// Run 单机版爬虫
func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request // 储存每一个任务，深度处理每一个页面下的 a 标签，直到没有

	//	初始化
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	//	逐个发起请求
	for len(requests) > 0 {
		request := requests[0]
		// 截取掉
		requests = requests[1:]

		fmt.Println()
		log.Printf("【请求URL】: %s\n", request.Url)
		fmt.Println()

		parseResult, err := worker(request)
		if err != nil {
			continue
		}

		// 收集
		requests = append(requests, parseResult.RequestArray...)
	}
}

func worker(r Request) (ParseRequest, error) {
	bytes, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("请求 %s 异常, err: %s \n", r.Url, err.Error())
		return ParseRequest{}, err
	}
	return r.ParseUrlFun(bytes), nil
}
