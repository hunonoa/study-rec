package hello02_data_enum

import "fmt"

/**
 * 枚举类型简单示例（注意：在Go里面一帮直接用const定义，也就是常量）
 */
func enums1() {
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	fmt.Println(cpp, java, python, golang)
}

/**
 * 注意：iota 表示0，也就是第一个，后面的会自动递增；
 * 也就是 java = 1，python = 2， golang = 3
 */
func enums2() {
	const (
		cpp = iota
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)
}
