package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type User struct {
	Name string
}

func main() {
	//getMethodTest1()
	//getMethodTest2()
	getZhengAi()
	// 对象转JSON String
	//str, err := json.Marshal(User)
	//var user User
	// JSON String转对象
	//err := json.Unmarshal(str, &user)
}

// 简单的get请求
func getMethodTest1() {
	res, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("获取信息出现错误:", err)
	}
	// 关闭请求
	defer res.Body.Close()
	// 注意：这个返回值是byte数组类型
	r, err := httputil.DumpResponse(res, true)
	if err != nil {
		fmt.Println("解码数据出现错误:", err)
	}
	fmt.Printf("发起请求获取到数据:%s", r)
}

func getMethodTest2() {
	// 创建请求体
	request, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	// 配置头信息
	request.Header.Add("token", "0xzxdsdsx1vsddas342dfgdgd55dfgdf5")
	// 发起请求
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("获取信息出现错误:", err)
	}
	// 关闭请求
	defer res.Body.Close()
	// 注意：这个返回值是byte数组类型
	r, err := httputil.DumpResponse(res, true)
	if err != nil {
		fmt.Println("解码数据出现错误:", err)
	}
	fmt.Printf("发起请求获取到数据:%s", r)
}

func getMethodTest3() {
	// 创建请求体
	request, err := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	// 配置头信息
	request.Header.Add("token", "0xzxdsdsx1vsddas342dfgdgd55dfgdf5")
	// 自定义Http Client
	client := http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		fmt.Println("服务器说要重定向")
		return nil
	}}
	// 发起请求
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("获取信息出现错误:", err)
	}
	// 关闭请求
	defer res.Body.Close()
	// 注意：这个返回值是byte数组类型
	r, err := httputil.DumpResponse(res, true)
	if err != nil {
		fmt.Println("解码数据出现错误:", err)
	}
	fmt.Printf("发起请求获取到数据:%s", r)
}

func getZhengAi() {
	request, err := http.NewRequest(http.MethodGet, "https://album.zhenai.com/u/1296109018", nil)
	if err != nil {
	}
	request.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36")
	request.Header.Set("cookie", "sid=1d309ed0-e032-477d-a913-16238d733077; ec=0RXUtRmV-1655345052308-fe998b768e46e1633930518; FSSBBIl1UgzbN7NO=5CSkbTZb739nCkxrKaa8jS4MSt7Jp1.4JhxnQftll4iYSiQINY4FrJhYJGKKlfK44Buc4oqoT0XmmL7CPkJ_g8q; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1655345068,1655774884,1655862257,1655952797; _efmdata=N6j5XcOz7TvWMhCDDeJTtDLOScwYkeavL4WSAFfQzLPwLd68m%2BmOusY4d1gRARw9u9jMIlfOICzBQmdX2nRFyKh07PcJRyW%2BtX8gFdJo8AE%3D; _exid=7ruM85QNopTpfn0LwLYTYwci413J8vyJzKxhp4XGb6wMIXFxXtptHwxM1XXy3PARR%2Bqz5QCchmlrWUdRr4Jiiw%3D%3D; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1655955831; FSSBBIl1UgzbN7NP=53.HfIKtugU3qqqDrOjNUNa9nHGvuGqvM4Cm8PyIlcBRm1UbKmaIBNoPq_1WZqKvEVzc3DIYLFuUJ_aAQt5.GXJ2G8RmQG8X00sd9taQALqzovDVdke3dvk62xa7LzWT4aoiq4KladJOFbwcRCgBjQAAR536wxhM1jcywQ5gSjmP7WKZj2brYl9kvSg307g.RV.JlUNt4rqTUd4M7BZO0JJUmW7LtA_afkBJMc1vHlGmeQ9RE3oIG3zzEbTiWU.UCa")
	// 自定义Http Client
	client := http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		fmt.Println("服务器说要重定向")
		return nil
	}}
	// 发起请求
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("获取信息出现错误:", err)
	}
	// 关闭请求
	defer res.Body.Close()
	// 注意：这个返回值是byte数组类型
	r, err := httputil.DumpResponse(res, true)
	if err != nil {
		fmt.Println("解码数据出现错误:", err)
	}
	fmt.Printf("发起请求获取到数据:%s", r)
}
