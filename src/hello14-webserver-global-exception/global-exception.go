package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	// 注意：添加下滑线该包没有被显示调用但是要引进来（pprof是性能分析程序，有了pprof之后http server程序可以使用 /debug/pprof 连接查看性能分析）
	_ "net/http/pprof"
	"os"
	"strings"
)

// 服务器实际调用的目标函数
func fileList(writer http.ResponseWriter, request *http.Request) error {
	// 获取到Path之后再截取字符串
	path := request.URL.Path[len("/src/"):]
	if strings.HasSuffix(path, "gogo") {
		var err CustomError = "访问连接不能以gogo结尾"
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("文件打开出现错误:", err)
		// 返回错误信息到前端
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		//fmt.Println("文件读取出现错误:", err)
		return err
	}
	writer.Write(all)
	return nil
}

// 定义一个类型它的本质是一个函数返回值是error
type appHandler func(writer http.ResponseWriter, request *http.Request) error

// 统一错误包装器函数，它的参数是目标函数，返回一个函数，服务器先调用返回的函数，再在返回的函数里面调用目标函数
func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			// 获取panic的参数
			r := recover()
			fmt.Println("程序内部出现错误，有人调用panic:", r)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}()
		// 调用目标函数
		err := handler(writer, request)
		if err != nil {
			httpCode := http.StatusOK
			// 如果错误是我们自定义的错误
			if userErr, ok := err.(userError); ok {
				fmt.Println("服务器出现自定义错误:", userErr)
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}
			switch {
			case os.IsNotExist(err):
				fmt.Println("服务器出现不存在的错误:", err)
				httpCode = http.StatusNotFound
			case os.IsPermission(err):
				fmt.Println("服务器出现没有权限错误")
				httpCode = http.StatusForbidden
			default:
				httpCode = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(httpCode), httpCode)
		}
	}
}

// 定义一个自定义错误接口，它包含error接口，就是说本生也是error接口,相当于继承了error接口（也就是要实现userError接口还要实现error接口）和一个Message函数
type userError interface {
	error
	Message() string
}

// 实现自定义错误接口
// 定义一个类型CustomError它的本质类型是string
type CustomError string

// 实现error接口
func (err CustomError) Error() string {
	return string(err)
}

// 实现userError接口
func (err CustomError) Message() string {
	return string(err)
}

// web server统一错误处理
// 可使用 http://127.0.0.1:8888/src/go.mod 连接测试以下代码
func main() {
	http.HandleFunc("/src/", errWrapper(fileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("服务器绑定出现错误:", err)
	}
}
