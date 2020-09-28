package scheduler

import "oliver/study/go-spider/engine"

type SimpleScheduler struct {
	workerChan chan engine.MyRequest
}

func (s *SimpleScheduler) WorkChan() chan engine.MyRequest {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(requests chan engine.MyRequest) {

}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.MyRequest)
}

func (s *SimpleScheduler) Submit(request engine.MyRequest) {
	// 放入 worker chan
	go func() {
		s.workerChan <- request
	}()
}
