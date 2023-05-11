package parser

import (
	"go-study/src/hello22-reptile/engine"
	"go-study/src/hello22-reptile/model"
	"regexp"
)

// 查找名字信息
var regName, _ = regexp.Compile(`<h1 data-v-cc1a17de="" class="nickName">(.+)</h1>`)

// 查找除名字以外其它的信息
var regOther, _ = regexp.Compile(`<div data-v-8b1eac0c="" class="m-btn purple">(.+)</div>`)

// 解析用户详细信息数据
func ParserUserProfile(contents []byte, name string) engine.ParserResult {
	// 匹配名字标签
	matchNames := regName.FindSubmatch(contents)
	matcheOthers := regOther.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	profile := model.Profile{}
	// 名字
	profile.Name = string(matchNames[1])
	for i, m := range matcheOthers {
		itemStr := string(m[1])
		switch i {
		// 婚姻状况
		case 0:
			profile.Marriage = itemStr
		// 年龄
		case 1:
			profile.Age = itemStr
		// 星座
		case 2:
			profile.Xinzuo = itemStr
		// 身高
		case 3:
			profile.Height = itemStr
		// 体重
		case 4:
			profile.Weight = itemStr
		// 户口
		case 5:
			profile.Hokou = itemStr
		// 收入
		case 6:
			profile.Income = itemStr
		// 职业
		case 7:
			profile.Occupation = itemStr
		// 学历
		case 8:
			profile.Education = itemStr
		case 9:
		case 10:
		case 11:
		}
	}
	result.Items = []interface{}{profile}
	return result
}
