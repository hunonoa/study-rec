package hello07_map

import "fmt"

// Map简单使用（注意：在Go里面除了数组，Map，Function之外其他类型都可以作为Map的Key，但是如果结构体里面包含数组，Map，Function属性也是不能作为Key的）
func mapTest() {
	// 中括号里面表示key的类型，后面跟着value的类型
	map1 := map[string]string{"name": "maomoa"}
	fmt.Println("", map1)
	// 往map里面添加元素
	map1["aaa"] = "asdasdas"
	// 获取map里面的元素
	aaaValue := map1["aaa"]
	// ok是boolean类型表示key是否存在
	aaaValue, ok := map1["aaa"]
	// 删除map里面的元素
	delete(map1, "aaa")
	fmt.Println("", aaaValue, ok)
	// 遍历map
	for key, value := range map1 {
		fmt.Println("", key, value)
	}
}

// 寻找最长不含有重复字符的子串
// 思路：
// 首先定义没有重复字符子串的开始位置，和最大长度
// 遍历整个字符串用map记录每个字符的的位置，在遍历的过程当中，用当前字符去map里面找是否有重复字符，如果有则调整开始位置，没有重复字符子串的最大长度等于当前遍历到的位置减开始位置
func test(s string) int {
	// 建立一个没有元素的map（key是rune类型（就是char字符类型），value是int类型）
	lastOccurred := make(map[rune]int)
	// 没有重复字符子串的开始位置
	start := 0
	//没有重复字符子串的最大长度
	maxLength := 0
	for i, ch := range []rune(s) {
		// 遇到重复字符调整子串的开始位置
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		// 调整没有重复字符子串的最大长度
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
