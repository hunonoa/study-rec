package main

import "fmt"

// 函数式编程和闭包函数简单使用

// 定义一个累加器函数它的返回值是一个函数（注意：这是一个闭包函数所以sum变量的状态会一直被保存）
func adder() func(value int) int {
	sum := 0
	return func(value int) int {
		sum += value
		return sum
	}
}

// 定义一个类型 iAdder它的实际类型是一个函数，它的参数int类型，返回值一个是int，一个iAdder
type iAdder func(int) (int, iAdder)

// 定义一个函数返回一个iAdder的闭包函数（注意：这是一个闭包函数所以base变量的状态会一直被保存）
func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

// 斐波那契数列简单实现（注意：这是一个闭包函数所以a，b变量的状态会一直被保存）
// 每一个数都是前两个数的和
func fibinacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 函数实现接口简单使用
// 定义一个类型Funcc它的实际类型是一个函数
type Funcc func() string

// 定义一个接口
type Maom interface {
	GetName() string
}

// 函数实现接口
func (funcc Funcc) GetName() string {
	return funcc()
}

func main() {
	adder := adder()
	// 0 + 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9
	for i := 0; i < 10; i++ {
		fmt.Println(adder(i))
	}

	iadder := adder2(1)
	sum, iadderr := iadder(1)
	fmt.Println("", sum, iadderr)

	fi := fibinacci()
	fmt.Println(fi())
	fmt.Println(fi())
	fmt.Println(fi())
	func() {
		fmt.Println("该函数体立马执行，因为最后加了括号")
	}()

}
