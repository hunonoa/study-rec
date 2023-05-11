package parser

import (
	"go-study/src/hello24-reptile-concurrent/engine"
	"regexp"
)

// 解析城市列表数据
func ParserCityList(contents []byte) engine.ParserResult {
	// ^ 代表取反；[^>]就是只只要不是>就可以匹配
	reg, _ := regexp.Compile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`)
	//matches := reg.FindAll(contents, -1)
	matches := reg.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for i, m := range matches {
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParserCityUser,
		})
		//fmt.Printf("Url:%s,City:%s\n", m[1], m[2])
		// 只要5个城市的数据
		if i > 1 {
			break
		}
	}
	return result
}
