package main

import "sync"

// sync同步锁简单使用

// 定义一个类型atomicInt它的原始类型是int
type atomicInt struct {
	value int
	// 锁
	lock sync.Mutex
}

// 递增
func (a *atomicInt) increment() {
	// 加锁
	a.lock.Lock()
	a.value++
	// 解锁（加了defer相当于在final里面解锁）
	defer a.lock.Unlock()
}

func (a *atomicInt) get() int {
	return a.value
}

func main() {

}
