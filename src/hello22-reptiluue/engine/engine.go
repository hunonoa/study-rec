package engine

import (
	"fmt"
	"go-study/src/hello22-reptile/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		// 截取数组
		requests = requests[1:]
		fmt.Println(r.Url)
		bytes, err := fetcher.Fetch(r.Url)
		if err != nil {
			fmt.Printf("%s\n", err)
			continue
		}
		parserResult := r.ParserFunc(bytes)
		// 合并数组；... 表示展开数组
		requests = append(requests, parserResult.Requests...)
		// 简单打印 Item
		for _, item := range parserResult.Items {
			fmt.Printf("%s,", item)
		}
		fmt.Println()
	}
}
