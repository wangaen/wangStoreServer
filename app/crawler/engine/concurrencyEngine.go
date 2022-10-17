package engine

import (
	"log"
)

// ConcurrencyEngine 并发版引擎
type ConcurrencyEngine struct {
	Scheduler
	WorkerCount int // 工人数
	//ItemChan    chan interface{}
}

// Scheduler 调度器
type Scheduler interface {
	Submit(Request)
	ConfigWorkerChan(chan Request)
	//Run()
	//WorkReady(chan Request)
}

func (c *ConcurrencyEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseRequest)
	c.Scheduler.ConfigWorkerChan(in)
	for i := 0; i < c.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, seed := range seeds {
		c.Scheduler.Submit(seed)
	}

	//	读取 out
	for {
		result := <-out
		for _, item := range result.TagContent {
			log.Println("遍历到，item : ", item)
		}
		for _, request := range result.RequestArray {
			c.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseRequest) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

//func (c *ConcurrencyEngine) Run(seeds ...Request) {
//	out := make(chan ParseRequest)
//
//	c.Scheduler.Run()
//
//	// 创建工人
//	for i := 0; i < c.WorkerCount; i++ {
//		CreateWorker(c.Scheduler.ConfigWorkerChan(), out, c.Scheduler)
//	}
//
//	// 种子任务分配调度器
//	for _, seed := range seeds {
//		c.Submit(seed)
//	}
//
//	//	处理 out
//	itemCount := 0
//	for {
//		fmt.Println("222222222222222")
//		result := <-out
//		for _, item := range result.TagContent {
//			//go func() { c.ItemChan <- item }()
//			log.Printf("itemCount: %d, %s", itemCount, item)
//			//if book, ok := item.(models.Book); ok {
//			//	book.PrintBookDetails()
//			//	book.WriteBookToFile()
//			//}
//			itemCount++
//		}
//		for _, request := range result.RequestArray {
//			c.Scheduler.Submit(request)
//		}
//	}
//}
//
//func CreateWorker(in chan Request, out chan ParseRequest, s Scheduler) {
//	go func() {
//		for {
//			s.WorkReady(in)
//			request := <-in                // 取出一个请求任务
//			result, err := worker(request) // 工人处理
//			if err != nil {
//				break
//			}
//			fmt.Println("44444444444", result)
//			out <- result // 处理完成写入
//		}
//	}()
//}
//
//func worker(request Request) (ParseRequest, error) {
//	log.Printf("请求 %s 中...", request.Url)
//	contentByte, err := fetcher.Fetch(request.Url)
//	if err != nil {
//		log.Printf("err: %v \n", err.Error())
//		return ParseRequest{}, err
//	}
//
//	result := request.ParseUrlFun(contentByte)
//
//	return result, nil
//}
