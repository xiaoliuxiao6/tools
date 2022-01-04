package tools

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Result struct {
	Response *http.Response // 原始返回信息
	Body     string         // 返回的 Body 字符串
	Duration time.Duration  // 执行耗时
}

func Request(url string, method string, data string, Headers map[string]string, parms map[string]string) (Resulter Result, err error) {
	startTime := time.Now() // 函数开始时间
	payload := strings.NewReader(data)

	client := &http.Client{}

	client.Timeout = 30 // 设置超时时间

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return
	}

	// 解析参数
	q := req.URL.Query()
	if len(parms) > 0 {
		for k, v := range parms {
			q.Add(k, v)
		}
	}
	req.URL.RawQuery = q.Encode()

	// 解析请求头
	req.Header.Add("Content-Type", "application/json")
	if len(Headers) > 0 {
		for k, v := range Headers {
			req.Header.Add(k, v)
		}
	}

	// 发送请求
	res, err := client.Do(req)
	if err != nil {
		return
	}

	Resulter.Response = res

	// 解析返回数据
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	Resulter.Body = string(body)

	defer res.Body.Close() // 延迟关闭

	Resulter.Duration = time.Since(startTime)
	return
}
