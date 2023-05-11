package parser

import (
	"go-study/src/hello22-reptile/engine"
	"regexp"
)

// 解析城市用户数据
func ParserCityUser(contents []byte) engine.ParserResult {
	// ^ 代表取反；[^>]就是只要不是>就可以匹配
	reg, _ := regexp.Compile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^<]*<img src="https://photo.zastatic.com/images/photo/[^"]+" alt="([^>]+)">[^<]*</a>`)
	//matches := reg.FindAll(contents, -1)
	matches := reg.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for i, m := range matches {
		var name string = string(m[2])
		result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(contents []byte) engine.ParserResult {

				return ParserUserProfile(contents, name)
			},
		})
		if i > 0 {
			break
		}
		//fmt.Printf("Url:%s,City:%s\n", m[1], m[2])
	}
	return result

}
