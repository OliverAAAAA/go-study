package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Items
	RequestProcessor Processor
}

type Processor func(MyRequest) (ParseResult, error)

type Scheduler interface {
	ReadyNotifyier
	Submit(MyRequest)
	WorkChan() chan MyRequest
	Run()
}

type ReadyNotifyier interface {
	WorkerReady(chan MyRequest)
}

func (c *ConcurrentEngine) Run(seeds ...MyRequest) {

	out := make(chan ParseResult)
	//初始化Scheduler的request管道
	c.Scheduler.Run()
	for i := 0; i < c.WorkerCount; i++ {
		//消费in，结果放入out
		c.createWorker(c.Scheduler.WorkChan(), out, c.Scheduler)

	}
	for _, r := range seeds {
		if idDuplicate(r.Url) {
			continue
		}
		c.Scheduler.Submit(r)
	}
	for {
		result := <-out
		for _, item := range result.Items {
			var data = item
			go func() { c.ItemChan <- data }()
		}
		for _, req := range result.Requests {
			if idDuplicate(req.Url) {
				continue
			}
			c.Scheduler.Submit(req)
		}
	}
}

func (c *ConcurrentEngine) createWorker(in chan MyRequest, out chan ParseResult, r ReadyNotifyier) {
	go func() {
		for {
			//tell schedule i am ready
			r.WorkerReady(in)
			//取request
			req := <-in
			//result,err := Worker(req)
			result, err := c.RequestProcessor(req)
			if err != nil {
				continue
			}
			//放结果
			out <- result
		}
	}()

}

var visitedUrls = make(map[string]bool)

func idDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
