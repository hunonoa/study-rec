package hello01

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

/**
 * 基础数据类型，变量定义，常量定义，数据类型强制转换
 */

// 定义常量（注意：Go里面常量名不一定要大写）
const aa = "sadasdsadas"

// 定义常量，同时定义多个常量
const (
	bbb = 10
	ccc = 20
)

func main() {
	// 注意：interface{}表示任意类型就相当于java的Object类型
	var object interface{} = 12
	fmt.Println("", object)
	// 注意：nil表示空指针;如果是结构体指定nil指针，是可以使用该指针调用结构体的函数的，它不会报错，但是取属性还是会报错
	var aa *string = nil
	fmt.Println(aa)
	// 定义常量
	const aaa string = "aaaaa"
	variableType()
	variableObject()
	convertObject()
	pointer()
}

// 字符串的相关操作
func stringTest() {
	str := "asdascxfds血常规撒大声地"
	// 分割字符串
	strings.Split(str, ",")
	// 去除首位两端的字符
	strings.Trim(str, "a")
	// 查找子字串是否存在
	strings.Contains(str, "as")
	// 转小写
	strings.ToLower(str)
	// 转大写
	strings.ToUpper(str)
	// 字符串转整数
	num, err := strconv.Atoi("3")
	fmt.Println("", num, err)
}

// 遍历打印一个字符串的每一个字符（注意：在Go里面没有char类型与之对应的是rune类型但是rune类型站4个字节）
func iteratorTest() {
	str := "我的测试程序sdfksdfksd"
	strCount := utf8.RuneCountInString(str)
	fmt.Println("字符串str有多少个字符", strCount)
	// 将str强制转换成byte数组
	bytes := []byte(str)
	for len(bytes) > 0 {
		// 第一个返回值是获取到一个byte数组里面的第一个字符，第二个返回值是第一个字符所站字节长度
		ch, size := utf8.DecodeRune(bytes)
		// 删除第一个字符的数据
		bytes = bytes[size:]
		// 打印第一个字符
		fmt.Println("", ch)
	}

	// 更简单的打印字符串的每一个字符（注意：这里是先将字符串转换成rune数组（就是char数组）再遍历rune数组，最后将字符一个一个拿出来）
	for i, char := range []rune(str) {
		fmt.Println("%i,%s", i, char)
	}
}

/**
 * 简单数据类型强制转换
 */
func convertObject() {
	var a1 int32 = 10
	// 强制转换
	var a2 float32 = float32(a1)
	// interface{} 表示任意类型
	var a3 interface{} = 12
	// 将任意类型进行强制转换
	var a4 = a3.(float32)
	fmt.Println(a1, a2, a4)
}

/**
 * 基础数据类型
 */
func variableType() {
	var abool bool = false
	var abyte byte = 'a'
	// 这个就是各大语言的char（字符）类型（注意：Go里面这个类型是占4个字节的，但是它只存一个字符）
	var arune rune = 'b'
	var afloat float32 = 10.01
	var acomplex complex64 = 10101
	fmt.Println(abool, abyte, arune, afloat, acomplex)
}

/**
 * 变量定义
 */
func variableObject() {
	// 定义跨行的字符串变量（注意：如果字符串变量要跨行那就必须使用``符号包起来）
	var straaa string = `
    asdasdasdas 
    asSSSDSDD
    `
	fmt.Println("", straaa)
	var a1 int32 = 1
	// 同时定义多个变量并赋值
	var a2, a3 int32 = 2, 3
	// 不指定类型直接定义变量
	var a6, a7 = "12545", 10
	var a4 string = "sfdasl"
	// 使用冒号定义变量，表示初始赋值，不需要写var
	a5 := 10
	fmt.Printf("a1=%d,a2=%d,a3=%d,a4=%q,a5=%d,a6=%q,a7=%d", a1, a2, a3, a4, a5, a6, a7)
}

func pointer() {
	var a int = 2
	// pa指针指向a变量内存地址
	var pa *int = &a
	// 修改pa指针内存地址的值，也就是a变量的值
	*pa = 3
	fmt.Println(a)
}

// 使用类型别名来定义变量（这个和Rust一样）

// 定义一个类型User它是实际类型string
type User string

// 定义一个类型AFunc它是实际类型是一个函数
type AFunc func(a int, b int) int32
