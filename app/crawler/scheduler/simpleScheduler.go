package scheduler

import "wangStoreServer/app/crawler/engine"

// 单机调度器

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}

func (s *SimpleScheduler) ConfigWorkerChan(cr chan engine.Request) {
	s.workerChan = cr
}
