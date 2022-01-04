package tools

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func Request(url string, method string, data string, Headers map[string]string, parms map[string]string) (body []byte, res *http.Response, err error) {

	payload := strings.NewReader(data)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return
		// log.Panicln(err)
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
	res, err = client.Do(req)
	if err != nil {
		return
		// log.Panicln(err)
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	defer res.Body.Close() // 延迟关闭

	return
}
