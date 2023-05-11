package main

import (
	"fmt"
	"sync"
)

// channel管道简单使用（往channel里面发送数据时必须要有空闲goroutine去接收数据，否则发送数据代码将会进入等待）

func channelDemo(wg *sync.WaitGroup) {
	// 创建一个channel它里面传输的数据类型是int（注意：chan表示就是channel）
	channel := make(chan int)
	// 定义只能发送数据的channel
	//channel1 := make(chan<- int)
	// 定义只能接收数据的channel
	//channel2 := make(<-chan int)
	// 定义一个goroutine一直从channel里面接收数据
	go work(channel, wg)
	// 往channel里面发送数据 1
	// 注意：往channel里面发送数据必须得先有一个goroutine去接收数据，否则代码将会进入死锁
	channel <- 1
	// 往channel发送数据 12
	channel <- 12
	// 关闭channel
	close(channel)
}

// 定义一个函数参数是一个channel管道传输的是int类型数据，还有一个参数是类似于CountDownLatch的东西（注意：该函数的作用是一直从channel里面去读取数据）
func work(channel chan int, wg *sync.WaitGroup) {
	// 死循环去接收数据
	for {
		// 从channel里面接收数据（注意：ok表示是否还有数据，也就是channel是否关闭）
		one, ok := <-channel
		if ok {
			fmt.Println("从channel里面接收到的数据是: ", one)
		} else {
			break
		}
	}
	// 减少计数器相当于java的countDown()函数
	wg.Done()
}

// 定义一个函数参数是一个channel管道传输的是int类型数据，还有一个参数是类似于CountDownLatch的东西（注意：该函数的作用是一直从channel里面去读取数据）
func work2(channel chan int, wg *sync.WaitGroup) {
	// 循环从channel里面去读取数据（注意：这个写法的好处是如果channel关闭了会自动跳出循环）
	for val := range channel {
		fmt.Println("从channel里面接收到的数据是: ", val)
	}
	// 减少计数器相当于java的countDown()函数
	wg.Done()
}

// 带缓存区的channel
func bufferChannelTest(wg *sync.WaitGroup) {
	// 创建一个带缓冲区的channel它里面传输的数据类型是int，缓冲区大小是3（简单理解队列大小是3）（注意：chan表示就是channel）
	channel := make(chan int, 3)
	go work2(channel, wg)
	channel <- 12
	channel <- 14
	channel <- 15
	// 关闭管道
	close(channel)
}

func main() {
	// 用于等待channel消费完成（这个相当于java的CountDownLatch）
	var wg sync.WaitGroup
	wg.Add(2)
	channelDemo(&wg)
	bufferChannelTest(&wg)
	// 等待channel消费完成
	wg.Wait()
}
