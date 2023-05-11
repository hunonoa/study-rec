package main

import "fmt"

// 接口组合简单使用（组合接口相当于继承其它接口）
// Go接口函数一般不多就两三个，如果函数多的话建议做拆分，再用接口组合的方式进行组合
// 定义第一个接口
type Dogo interface {
	action()
}

// 定义第二个接口
type Godo interface {
	pop()
}

// 定义组合接口（就是将多个接口组合到一起，到时候当参数传递的时候，就必须传递一个实现了多个接口的实体）
// 注意：组合接口表示本身也是其它接口，就是说Anmi接口本身也是Dogo和Godo接口
type Anmi interface {
	Dogo
	Godo
}

// 定义一个结构体用来实现多个接口
type AnmiImpl struct {
}

// 实现Godo接口的pop()函数
func (ai AnmiImpl) pop() {
	fmt.Println("调用了pop()函数")
}

// 实现Dogo接口的action()函数
func (ai AnmiImpl) action() {
	fmt.Println("调用了action()函数")
}

// 传递一个组合接口的参数就可以调用多个接口的函数（注意：组合接口其实表示的就是多个接口）
func TestInterface(anmi Anmi) {
	anmi.pop()
	anmi.action()
}
