package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// 错误处理简单使用
// panic() 会停止函数执行，当然是在defer之后
// recover() 可以获取到调用panic()函数所传的参数

// defer关键字相当于JAVA的Final在函数返回或结束时执行（defer管理一个栈空间，先进后出，也就是说多个defer代码，写在前面的后执行，写在后面的先执行）
func tryFinal() {
	defer fmt.Println("函数返回执行1")
	defer fmt.Println("函数返回执行2")
	fmt.Println("函数执行结束")
}

// 注意：defer语句如果有参数那么参数的状态会在执行时被保留也就是说下面的代码会以此打印 4,3,2,1,0
func tryFinal2() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// 错误处理
func writeFile(fileName string) {
	//file, err := os.Create(fileName)
	file, err := os.OpenFile(fileName, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		//panic(err)
		fmt.Println("程序出错了:", err.Error())
		// 错误如果是PathError
		if pathError, ok := err.(*os.PathError); !ok {
			fmt.Println("程序不是出现路径错误:", pathError)
		} else {
			fmt.Println("程序出现路径错误", pathError.Op, pathError.Path, pathError.Err)
		}
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 10; i++ {
		writer.WriteString(string(i))
	}
}

// 自定义错误
func customError() {
	errors.New("自定义错误")
}

// panic和recover简单使用
func panicAndrecover() {
	defer func() {
		// 获取panic的参数
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("程序调用panic停止函数执行，panic的参数是error类型:", err)
		} else {
			fmt.Println("程序调用panic停止函数执行，panic的参数不是error类型")
		}
	}()
	b := 0
	a := 5 / b
	fmt.Println(a)
}

func main() {
	tryFinal()
	tryFinal2()
	writeFile("test.file")
	panicAndrecover()
	func() {
		fmt.Println("该函数体立马执行，因为最后加了括号")
	}()
}
