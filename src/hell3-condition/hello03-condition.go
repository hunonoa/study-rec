package hello03_condition

import "reflect"
import "runtime"
import "fmt"

/**
 * 条件，分支，循环，switch语句简单使用以及函数定义反射简单使用
 */

/**
 * 定义函数，它的参数是count,返回值类型是string
 */
func fnTest(count int32) string {
	// 接收返回多个值的函数
	a1, a2 := fntestTest()
	// _ 表示不使用该返回值
	a3, _ := fntestTest()
	fmt.Println(a1, a2, a3)
	return "sadasda"
}

/**
 * 函数返回多个值
 */
func fntestTest() (string, uint32) {
	return "", 10
}

/**
 * 定义函数返回指定变量的值（注意：下面的函数就是返回a1和a2在函数里面最后的值）
 */
func fnTestTest() (a1, a2 uint32) {
	// 调用函数参数带有函数的方法
	fnParamTest(func(a1, a2 int) int {
		return a1 - a2
	}, 1, 10)
	a1 = 10
	a2 = 20
	return
}

/**
 * 函数传入函数参数（注意：fn参数表示一个函数且有返回值int类型）
 */
func fnParamTest(fn func(int, int) int, a, b int) int {
	// 使用反射获取函数指针
	f := reflect.ValueOf(fn).Pointer()
	// 获取函数名称
	fName := runtime.FuncForPC(f).Name()
	fmt.Println(fName)
	return fn(a, b)
}

func condition() {
	var count uint32 = 10
	if count > 10 {
		fmt.Println("大于10")
	} else {
		fmt.Println("没有大于10")
	}

}

/**
 * switch语句简单使用（注意：不需要写break，它会自动break）
 */
func switchTest() {
	var ca string = "="
	var res uint32
	switch ca {
	case "+":
		res = 10
	default:
		res = 100
	}
	fmt.Println(res)
}

/**
 * 循环测试
 */
func forTest() {
	sum := 1
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)
	// Go里面没有while所以for可以这样写
	for true {

	}
	// 自旋循环
	for {

	}
}
