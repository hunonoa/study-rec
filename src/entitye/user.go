package entity

// 定义一个结构体（注意：为结构体定义的函数必须在同一个包内，但是可以在不同的文件内）（注意：首字母大写表示 public，首字母小写表示 private）
type User struct {
	Name string
	Age  int
}
