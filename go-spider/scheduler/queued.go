package scheduler

import "oliver/study/go-spider/engine"

type QueuedScheduler struct {
	requestChan chan engine.MyRequest
	workerChan  chan (chan engine.MyRequest)
}

func (s *QueuedScheduler) Submit(r engine.MyRequest) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.MyRequest) {
	s.workerChan <- w
}

func (s *QueuedScheduler) WorkChan() chan engine.MyRequest{
	return make(chan engine.MyRequest)
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.MyRequest)
	s.requestChan = make(chan engine.MyRequest)
	go func() {
		var requestQ []engine.MyRequest
		var workerQ []chan engine.MyRequest

		for {
			var activeRequest engine.MyRequest
			var activeWorker chan engine.MyRequest
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)

			case w := <-s.workerChan:
				workerQ = append(workerQ, w)

			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}

	}()
}
