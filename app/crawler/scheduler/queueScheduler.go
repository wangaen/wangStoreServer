package scheduler

import "wangStoreServer/app/crawler/engine"

// 队列调度器

type QueueScheduler struct {
	RequestChan chan engine.Request
	WorkerChan  chan chan engine.Request
}

func (q *QueueScheduler) Submit(request engine.Request) {
	q.RequestChan <- request
}

//func (q *QueueScheduler) ConfigWorkerChan(requests chan engine.Request) {
//	panic("implement me")
//}

func (q *QueueScheduler) WorkReady(w chan engine.Request) {
	q.WorkerChan <- w
}

func (q *QueueScheduler) Run() {
	q.WorkerChan = make(chan chan engine.Request)
	q.RequestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWork chan engine.Request

			if len(requestQ) > 0 && len(workQ) > 0 {
				activeRequest = requestQ[0]
				activeWork = workQ[0]
			}

			select {
			case r := <-q.RequestChan:
				requestQ = append(requestQ, r)
			case w := <-q.WorkerChan:
				workQ = append(workQ, w)
			case activeWork <- activeRequest:
				workQ = workQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
