package engine

import (
	"log"
	"wangStoreServer/app/crawler/fetcher"
	"wangStoreServer/app/crawler/models"
)

type Scheduler interface {
	Submit(Request)
	ConfigWorkerChan() chan Request
	Run()
	WorkReady(chan Request)
}

type ConcurrencyEngine struct {
	Scheduler
	WorkerCount int // 工人数
}

func (c *ConcurrencyEngine) Run(seeds ...Request) {

	out := make(chan ParseRequest)

	c.Scheduler.Run()

	// 创建工人
	for i := 0; i < c.WorkerCount; i++ {
		CreateWorker(c.Scheduler.ConfigWorkerChan(), out, c.Scheduler)
	}

	// 种子任务分配调度器
	for _, seed := range seeds {
		c.Submit(seed)
	}

	//	处理 out
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.TagContent {
			log.Printf("itemCount: %d, %s", itemCount, item)
			if book, ok := item.(models.Book); ok {
				book.PrintBookDetails()
				book.WriteBookToFile()
			}
			itemCount++
		}
		for _, request := range result.RequestArray {
			c.Scheduler.Submit(request)
		}
	}
}

func CreateWorker(in chan Request, out chan ParseRequest, s Scheduler) {
	go func() {
		for {
			s.WorkReady(in)
			request := <-in                // 取出一个请求任务
			result, err := worker(request) // 工人处理
			if err != nil {
				continue
			}
			out <- result // 处理完成写入
		}
	}()
}

func worker(request Request) (ParseRequest, error) {
	log.Printf("请求 %s 中...", request.Url)
	contentByte, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("请求 %s 异常, err: %s \n", request.Url, err.Error())
		return ParseRequest{}, err
	}

	result := request.ParseUrlFun(contentByte)

	return result, nil
}
