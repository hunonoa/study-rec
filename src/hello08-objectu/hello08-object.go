package hello08_object

import "fmt"

// 面向对象简单使用，Go里面没有继承和多态（注意：首字母大写表示 public，首字母小写表示 private）

// 定义一个结构体（注意：为结构体定义的函数必须在同一个包内，但是可以在不同的文件内）
type TreeNode struct {
	// 属性value类型是int
	value int
	// 属性left和right类型都是 TreeNode指针
	left, right *TreeNode
}

// 为结构体TreeNode定义一个名字叫 print 的函数（注意：这样的函数就可以使用 结构体变量加点调用，比如: treeNode.print() ）
// 注意：如果不加*默认传过来的是值那么就不能修改原结构体的值，因为Go默认都是值传递
func (node *TreeNode) print() {
	fmt.Println("", node)
}

// 为结构体定义setValue函数
// 注意：如果不加*默认传过来的是值那么就不能修改原结构体的值，因为Go默认都是值传递
func (node *TreeNode) setValue(value int) {
	node.value = value
}

// 为结构体实现toString()函数
func (node *TreeNode) String() string {
	return fmt.Sprintf("TreeNode:{value=%d}", node.value)
}
func createTreeNode(value int) *TreeNode {
	return &TreeNode{value: value}
}

func objectTest() {
	// 初始化某个结构体
	rootNode := TreeNode{value: 1}
	// 定义一个结构体变量
	var treeNode TreeNode
	// 初始化结构体（注意：nil表示空指针，但是它安全的也可以调用函数,不会报错，但是如果取属性还是会报错的）
	treeNode = TreeNode{
		value: 3,
		left:  nil,
		right: nil,
	}
	// 为结构体的某个属性赋值
	rootNode.left = &treeNode
	// 定义结构体数组
	nodes := []TreeNode{{value: 1}, {value: 60}}
	fmt.Println("", treeNode, rootNode, nodes)
}
