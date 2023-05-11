package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// select简单使用（注意：该章节包含很多channel内容，关于channel内容请看上一章节）
// select使用case语法处理channel里面的数据；或接收或发送（注意：那个管道的数据先到就先处理谁）

func generator(wg *sync.WaitGroup) chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(6000)) * time.Millisecond)
			// 往管道发送数据
			out <- i
			i++
		}
		// 减少计数器相当于java的countDown()函数
		wg.Done()
	}()
	return out
}

func main() {
	// 用于等待channel消费完成（这个相当于java的CountDownLatch）
	var wg sync.WaitGroup
	wg.Add(2)
	var c1, c2 chan int = generator(&wg), generator(&wg)
	input := make(chan int)
	go func() {
		for {
			// 从channel里面接收数据（注意：ok表示是否还有数据，也就是channel是否关闭）
			one, ok := <-input
			if ok {
				fmt.Println("收到了input管道里面的数据:", one)
			}
		}
	}()
	var values []int
	// 该函数返回一个channel并在3秒后往该channel里面写一个时间数据
	afterChan := time.After(3 * time.Second)
	// 该函数返回一个channel每1秒都往channel里面写一个时间数据
	tick := time.Tick(time.Second)
	for {
		var activeWork chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWork = input
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
			fmt.Println("c1管道拿到了数据:", n)
		case n := <-c2:
			values = append(values, n)
			fmt.Println("c2管道拿到了数据:", n)
		case activeWork <- activeValue:
			values = values[1:]
			fmt.Println("将activeValue写入管道activeWork")
		// 该函数返回一个channel并在800毫秒后往该channel里面写一个时间数据（注意：如果select执行到这里来了就是说明select在800毫秒内没有执行到其他case操作）
		case <-time.After(800 * time.Millisecond):
			fmt.Println("800毫秒内select没有获取到管道里面数据")
		// 获取afterChan管道里面的数据
		case <-afterChan:
			fmt.Println("values数组的长度:", len(values))
			fmt.Println("3秒后收到afterChan管道里面的数据")
		// 获取tick管道里面的数据
		case <-tick:
			fmt.Println("每1秒获取一次数据，注意：tick管道的数据生产者是每1秒发送一次数据")
		// 注意：default一般不用
		default:
			//fmt.Println("没有从管道里面拿到数据")
		}
	}
	wg.Wait()
}
