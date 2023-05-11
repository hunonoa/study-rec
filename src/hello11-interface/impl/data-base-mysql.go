package impl

// 实现接口
// 先定义一个结构体
type MySql struct {
}

// 在结构体MySql里面定义接口DataBase里面所有的函数，就表示实现接口DataBase里面的函数
func (mysql *MySql) GetName() string {
	return "MySql"
}
