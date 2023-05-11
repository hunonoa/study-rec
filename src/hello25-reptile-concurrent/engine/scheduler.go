package engine

type QueuedScheduler struct {
	// 该管道用于Request入队（就是Request会先进入该管道然后会有程序从该管道里面获取Request最后将Request放入队列）
	requestChan chan Request
	// 该管道用于Work管道入队（就是Work管道会先进入该管道然后会有程序从该管道里面获取Work管道最后将Work管道放入队列）
	workChan chan chan Request
}

// 将Request放入入队管道
func (q *QueuedScheduler) Submit(r Request) {
	q.requestChan <- r
}

// 将Work管道放入入队管道
func (q *QueuedScheduler) WorkerReady(w chan Request) {
	q.workChan <- w
}

// 启动Scheduler调度器
func (q *QueuedScheduler) Run() {
	q.workChan = make(chan chan Request)
	q.requestChan = make(chan Request)
	go func() {
		// 定义请求队列
		var requestQ []Request
		// 定义Work管道队列
		var workQ []chan Request
		for {
			var activeRequest Request
			var activeWork chan Request
			// 如果Request队列和Work管道队列里面都有值（表示可以将Request放入Work管道）
			if len(requestQ) > 0 && len(workQ) > 0 {
				activeRequest = requestQ[0]
				activeWork = workQ[0]
			}
			select {
			// 获取Request入队
			case r := <-q.requestChan:
				requestQ = append(requestQ, r)
			//	获取Work管道入队
			case w := <-q.workChan:
				workQ = append(workQ, w)
			// 如果Request队列和Work队列里面都有值就将 activeRequest 发送到 activeWork通道
			case activeWork <- activeRequest:
				// 如果数据发送成功，就将该数据从队列里面拿掉
				workQ = workQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
