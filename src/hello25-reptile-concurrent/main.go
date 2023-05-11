package main

import (
	"go-study/src/hello25-reptile-concurrent/engine"
	"go-study/src/hello25-reptile-concurrent/zhenai/parser"
)

// 并发进阶版爬虫（使用队列来装请求和Work管道）
func main() {
	//FetchCity()
	request := engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	}
	simpleScheduler := engine.QueuedScheduler{}
	engine := engine.ConcurrentEngine{
		Scheduler: &simpleScheduler,
		WorkCount: 5,
	}
	engine.Run(request)
}
