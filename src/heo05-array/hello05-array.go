package hello05_array

import "fmt"

// 数组简单使用（数组切片请看下一章）

func testArray() {
	// 定义5个长度的数组（注意：没有赋值的数组默认每个元素的值都是0）
	var a1 [5]int
	// 指定数组长度
	a2 := [3]int{1, 2, 3}
	// 自动识别数组长度
	a3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a1, a2, a3)
	// 定义一个二维数组
	var a4 [5][4]int
	fmt.Println(a4)
	// 遍历数组
	for i := 0; i < len(a3); i++ {
		fmt.Println(a3[i])
	}
	// i 为下标
	for i := range a3 {
		fmt.Println(i)
	}
	// i为下标，v 为value
	for i, v := range a3 {
		fmt.Println(i, v)
	}
	// _表示不使用该值（这个值实际是下标），i为value
	for _, v := range a3 {
		fmt.Println(v)
	}
}

// 注意：参数是一定是一个长度为5的int类型数组的值（注意：如果参数没有标明数组长度那么参数就是一个数组的切片）
func printArray(array [5]int) {
	for i := range array {
		fmt.Println(i)
	}
}
