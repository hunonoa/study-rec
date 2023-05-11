package main

import (
	"fmt"
	"regexp"
)

// 正则表达式简单使用

const text = "My email is jiangjian@outlook.com"

func main() {
	// ^ 代表取反；[^>]就是只只要不是>就可以匹配
	// 编译正则表达式（双斜杠表示转义）
	reg, err := regexp.Compile("[a-z-A-Z-0-9]+@.+\\.(com)")
	// 编译正则表达式不会出现错误
	//reg := regexp.MustCompile("jiangjian@outlook.com")
	// 用反引号包起来的字符串可以使用单引号转义
	//reg, err := regexp.Compile(`.+@.+\.com`)
	if err != nil {
	}
	// 匹配字符串（注意：这个是只查找第一个匹配）
	match := reg.FindString(text)
	// 匹配字符串（注意：这个会查找所有字符串）
	matchs := reg.FindAllString(text, -1)
	// 匹配字符串以及子匹配（子匹配就是正则表达式里面带括号的部分）
	sunmatch := reg.FindStringSubmatch(text)
	fmt.Println("%s", match, matchs, sunmatch)
}
