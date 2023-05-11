package hello09_package

// 注意：同一目录下的文件包名必须一致（注意是文件，如果是子目录就可以取另一个包名）
import (
	"fmt"
	// 引入自定义包
	"go-study/src/entity"
)

// 自定义包简单使用和扩展已有类型的函数（注意：同一个目录下的文件，只能是一个包名）
// （注意：首字母大写表示 public，首字母小写表示 private）
func testUser() {
	user := entity.User{Name: "maoma", Age: 22}
	fmt.Println("", user)
}

// 扩展已有类型函数，第一种方法：通过一个新的结构体包含原有结构体指针来扩展原有结构体函数
type MyUser struct {
	user *entity.User
}

// 为结构体MyUser定义一个名字叫 getName 的函数（注意：这样的函数就可以使用 结构体变量加点调用，比如: myUser.getName() ）
func (myUser *MyUser) getName() string {
	if myUser != nil {
		return myUser.user.Name
	}
	return ""
}

// 扩展已有类型函数，第二种方法：通过别名的方法来扩展来原有结构体函数
type UserMy entity.User

func (userMy *UserMy) getName() string {
	if userMy != nil {
		return userMy.Name
	}
	return ""
}
