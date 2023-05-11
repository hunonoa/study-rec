package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	// 每一100毫秒会从管道里面获取一个数据（其他时间处于等待状态，间接实现Thread.Sleep操作）
	<-rateLimiter
	// 创建请求体
	request, err := http.NewRequest(http.MethodGet, url, nil)
	// 配置头信息
	request.Header.Add("referer", "http://www.zhenai.com/")
	request.Header.Add("sec-ch-ua-platform", "Linux")
	request.Header.Add("sec-ch-ua", " Not A;Brand\";v=\"99\", \"Chromium\";v=\"102\", \"Google Chrome\";v=\"102")
	request.Header.Add("sec-ch-ua-mobile", "?0")
	request.Header.Add("sec-fetch-dest", "document")
	request.Header.Add("sec-fetch-mode", "navigate")
	request.Header.Add("sec-fetch-site", "none")
	request.Header.Add("sec-fetch-user", "?1")
	request.Header.Add("pragma", "no-cache")
	request.Header.Add("upgrade-insecure-requests", "1")
	request.Header.Add("cache-control", "no-cache")
	request.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	//request.Header.Add("accept-encoding", "gzip, deflate, br")
	request.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36")
	request.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Add("cookie", "sid=1d309ed0-e032-477d-a913-16238d733077; ec=0RXUtRmV-1655345052308-fe998b768e46e1633930518; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1655345068,1655774884; FSSBBIl1UgzbN7NO=5CSkbTZb739nCkxrKaa8jS4MSt7Jp1.4JhxnQftll4iYSiQINY4FrJhYJGKKlfK44Buc4oqoT0XmmL7CPkJ_g8q; _exid=JRvLvl%2FfqsandO8QP21TEc9Y7f75Un5qdw75hT2p%2FEjuXVDMH4KL7lt4iw7otK7P4UQkQokpVA4RFAeXYUi%2FFA%3D%3D; _efmdata=N6j5XcOz7TvWMhCDDeJTtDLOScwYkeavL4WSAFfQzLPwLd68m%2BmOusY4d1gRARw9fq5XjvM5ld5fXJRVo2KbACDFM%2B6G%2Bpv4Jd1Tt%2BEEECo%3D; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1655806270; FSSBBIl1UgzbN7NP=53.DRpDtSbs0qqqDrfQkelaiTpxlVngvSnx_f64_M23KouNhba4BVRSCjhAaZEblRWTbcbHiuw9Vg_59DQpBCs0rjMfq9VRPGqTb0HddwQHvkufemrM9PbMTGh5xkburKGg5Y_VgjoS2ztFiKT_CE5M3U.ARYqFC6T5nbBiwB4QpY4uTomK72XUnV8UQNAajm.YjSfwPj_nS2Q.WHSRkNNMIiqMPz646P0K_WrKEg6xS1bGaXSB0CNNeeIPO0RTGeHMPqWG3BwET7kjFMPp4QOG")
	// 发起请求
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		// 判断数据编码格式
		e := DetermineEncoding(res.Body)
		// 转码（将GBK转码成UTF8）
		utf8Reader := transform.NewReader(res.Body /*simplifiedchinese.GBK.NewDecoder()*/, e.NewDecoder())
		all, err := ioutil.ReadAll(utf8Reader)
		if err != nil {
			return nil, err
		}

		return all, err
	}
	return nil, fmt.Errorf("请求%s失败,原因:%s", url, res)
}

// 判断数据编码格式
func DetermineEncoding(reader io.Reader) encoding.Encoding {
	// 读取数据流前1024个字节
	bytes, err := bufio.NewReader(reader).Peek(1024)
	if err != nil {
		fmt.Println("判断编码格式异常:", err)
		// 返回默认编码格式
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
