package tools

import (
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Request(url string, method string, data string, Headers map[string]string, parms map[string]string) (string, *http.Response) {

	payload := strings.NewReader(data)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Panicln(err)
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
		log.Panicln(err)

	}
	defer res.Body.Close() // 延迟关闭

	// fmt.Println("状态码为：", res.StatusCode)
	if res.StatusCode == 429 {
		log.Panicln("被限速，退出")
	}

	// 判断是有错误并返回
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panicln(err)
	}
	return string(body), res
}
