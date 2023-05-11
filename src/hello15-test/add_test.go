package hello15

import "testing"

// 注意：测试代码文件名以 _test 结尾

// 注意：简单测试函数的入口的参数必须是 t *testing.T
func TestAdd1(t *testing.T) {
	// 定义测试数据（注意：测试数据是一个结构体数组，结构体里面有3个属性分别是a,b,c类型是int32
	tests := []struct {
		a, b, c int32
	}{
		{1, 2, 3},
		{1, 2, 4},
	}

	for _, tt := range tests {
		if val := add(tt.a, tt.b); val != tt.c {
			t.Errorf("add(%d,%d) != %d", tt.a, tt.b, tt.c)
		}
	}
}

// 注意：性能测试（就是测试某个函数的性能，它会在函数后面显示执行时间）函数的入口的参数必须是 b *testing.B
// 下面的代码编译报错
/*func TestAddAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(1, 2)
	}
}*/
