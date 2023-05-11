package main

import (
	"bufio"
	"fmt"
	"go-study/src/hello22-reptile/engine"
	"go-study/src/hello22-reptile/zhenai/parser"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

// 单机版爬虫
func main() {
	//FetchCity()
	request := engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	}
	engine.Run(request)
}

func FetchCity() string {
	res, err := http.Get("https://www.zhenai.com/zhenghun")
	if err != nil {

	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		// 判断数据编码格式
		e := DetermineEncoding(res.Body)
		// 转码（将GBK转码成UTF8）
		utf8Reader := transform.NewReader(res.Body /*simplifiedchinese.GBK.NewDecoder()*/, e.NewDecoder())
		all, err := ioutil.ReadAll(utf8Reader)
		if err != nil {

		}
		extractCityList(all)
		//fmt.Printf("%s", all)
	}
	return ""
}

// 判断数据编码格式
func DetermineEncoding(reader io.Reader) encoding.Encoding {
	// 读取数据流前1024个字节
	bytes, err := bufio.NewReader(reader).Peek(1024)
	if err != nil {

	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

// 提取城市列表
func extractCityList(contents []byte) {
	// ^ 代表取反；[^>]就是只只要不是>就可以匹配
	reg, _ := regexp.Compile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`)
	//matches := reg.FindAll(contents, -1)
	matches := reg.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("Url:%s,City:%s\n", m[1], m[2])
	}
}
