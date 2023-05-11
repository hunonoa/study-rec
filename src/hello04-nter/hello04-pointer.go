package hello04_pointer

import "fmt"

/**
 * 指针简单使用
 */

func pointerr() {
	var a int = 2
	// pa指针指向a变量内存地址
	var pa *int = &a
	// 修改pa指针内存地址的值，也就是a变量的值
	*pa = 3
	fmt.Println(a)
	// 将指针传递给函数
	fnPointer(&a)
}

/**
 * 函数接收指针参数
 */
func fnPointer(a *int) {

}

/**
 * 交换两个参数的值（注意：参数是指针，交换两个指针地址即可交换两个参数的值）
 */
func swap(a, b *int) {
	*a, *b = *b, *a
}
