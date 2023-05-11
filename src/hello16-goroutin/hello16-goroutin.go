package main

import (
	"fmt"
	"runtime"
	"time"
)

// goroutine 并发简单使用
// 协程是非抢占式多任务处理（线程控制权由CPU调度，属于抢占式多任务处理），控制权由协程主动交出控制权，也就是说协程属于虚拟机层面的多任务
// 单个协程可以一个或多个线程上执行
func main() {
	for i := 0; i < 10; i++ {
		// 并发执行一个函数，该函数的参数是 i（注意：函数可以任意定义）
		go func(i int) {
			fmt.Println(i)
			// 让协程交出控制权（相当于线程的 yield() 函数）
			runtime.Gosched()
		}(i) // 传递参数 i
	}
	time.Sleep(time.Millisecond)
}
