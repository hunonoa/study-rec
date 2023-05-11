package main

import (
	"fmt"
	"go-study/src/hello11-interface/impl"
)

// Go接口定义简单使用（注意：Go里面实现接口其实是只要和接口有相同的函数就表示实现了接口）

/**
 * 定义一个接口（注意：Go里面实现接口其实是只要结构体对象里面包含接口函数就等于实现了接口）
 * Go接口函数一般不多就两三个，如果函数多的话建议做拆分，再用接口组合的方式进行组合
 */
type DataBase interface {
	// 定义一个接口函数
	GetName() string
}

// 测试接口当做参数传递（注意：在该函数里面只用到了接口的GetName()函数，所以传给该函数的参数只要是包含GetName()函数的就可以）
func TestTes(dataBase DataBase) {
	// 判断接口实际类型
	if dataBaseMySql, ok := dataBase.(*impl.MySql); ok {
		fmt.Println("传过来的DataBase接口实际类型是MySql：", dataBaseMySql)
	}
	if dataBaseOracle, ok := dataBase.(impl.Oracle); ok {
		fmt.Println("传过来的DataBase接口实际类型是Oracle：", dataBaseOracle)
	}

	// 第二中判断类型的方法
	switch m := dataBase.(type) {
	case impl.Oracle:
		fmt.Println("传过来DataBase接口的实现是Oracle", m)
	// *号表示case指针
	case *impl.MySql:
		fmt.Println("传过来DataBase接口的实现是MySql", m)
	}
	dataBaseName := dataBase.GetName()
	fmt.Println("", dataBaseName)
}

// 接口简单测试
func main() {

	var dataBase DataBase
	dataBase = impl.Oracle{}
	//dataBase = &impl.MySql{}

	TestTes(dataBase)
	fmt.Println("Hello")

	// 测试组合接口简单使用（就是如果一个函数的参数需要实现多个接口，那么我们就定义一个组合接口将多个接口组合到一起）
	var anmi = AnmiImpl{}
	TestInterface(anmi)
}
