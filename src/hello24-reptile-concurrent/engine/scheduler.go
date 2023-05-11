package engine

type SimpleScheduler struct {
	workChan chan Request
}

func (s SimpleScheduler) Submit(r Request) {
	// 将请求发送到数据通道
	go func() {
		s.workChan <- r
	}()
}

func (s *SimpleScheduler) ConfigureWorkChan(in chan Request) {
	s.workChan = in
}
