package engine

import (
	"fmt"
	"go-study/src/hello24-reptile-concurrent/fetcher"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureWorkChan(chan Request)
}

func (e ConcurrentEngine) Run(seeds ...Request) {
	// 创建请求发送通道
	in := make(chan Request)
	// 创建数据解析结果通道
	out := make(chan ParserResult)
	// 将数据发送通道给Scheduler
	e.Scheduler.ConfigureWorkChan(in)

	// 创建Work
	for i := 0; i < e.WorkCount; i++ {
		createWork(in, out)
	}
	// 将种子页面送给Work去解析
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	for {
		// 从数据解析结果通道获取数据
		parserResult := <-out
		// 简单打印解析结果 Item
		for _, item := range parserResult.Items {
			fmt.Printf("%s,", item)
		}
		fmt.Println()
		// 将子页面的解析再次送给work让其取解析
		for _, request := range parserResult.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWork(in chan Request, out chan ParserResult) {
	go func() {
		// 不断从请求发送通道里面获取要发送的请求
		for {
			request := <-in
			parserResult, err := work(request)
			if err == nil {
				// 将数据解析结果发送到通道
				out <- parserResult
			}
		}
	}()
}

func work(r Request) (ParserResult, error) {
	bytes, err := fetcher.Fetch(r.Url)
	if err != nil {
		fmt.Printf("请求地址: %s，发生异常: %s\n", r.Url, err)
		return ParserResult{}, err
	}
	return r.ParserFunc(bytes), nil
}
