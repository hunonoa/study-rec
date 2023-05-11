package parser

import (
	"fmt"
	"go-study/src/hello24-reptile-concurrent/fetcher"
	"testing"
)

func TestParserCityList(t *testing.T) {
	contents, err := fetcher.Fetch("https://www.zhenai.com/zhenghun")
	if err != nil {
	}
	result := ParserCityList(contents)
	fmt.Println(result)
}
