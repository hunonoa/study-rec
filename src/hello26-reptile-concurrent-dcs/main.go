package main

import (
	"go-study/src/hello26-reptile-concurrent-dcs/rpx"
	"log"
)

// 分布式爬虫简单架构
/**
抽象出3个服务
爬虫服务端：控制程序用于分发爬虫任务和接收任务结果
爬虫任务端：接收爬虫任务处理任务，完成后将任务结果发送给爬虫服务端
爬虫结果处理端：处理来自爬虫服务端发送过来的爬虫结果

爬虫结果处理端  <->   爬虫服务端   <->   爬虫任务端

*/

func main() {
	err := rpx.ServeRpc("127.0.0.1:1234", rpx.DemoService{})
	if err != nil {
		log.Printf("JSONRPC启动失败:%s", err)
	}
}
