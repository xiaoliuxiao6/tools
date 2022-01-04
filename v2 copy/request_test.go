package tools_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/xiaoliuxiao6/tools/v2"
)

// 基本使用示例
func TestCookie(t *testing.T) {
	url := "https://www.hpool.in/api/pool/list?type=opened"
	method := "POST"
	data := ""
	headers := map[string]string{
		"cookie": "_ga=GA1.2.1244375557.1639365528; lang=zh; MEIQIA_TRACK_ID=1ucBijMktnNZkNfDEUvhI2fuG69; auth_token=eyJldCI6MTY0MTQ0OTIzNiwiZ2lkIjo2NCwidWlkIjoyMjQ2MjI0LCJ2IjowfQ==.kndo3Lx4kcxdDbi4x7Y0u5wC4x1dLomM9Utz2y/K8PFZlLx87dGy26bJ0DtXipji/wTSEHsyTa5q9ZgA2KhbHrs7oLxHZ+deJiDo8OSgox+CYaEjff+0IjGCgYCLeJUw+GiH+xcNY1qaCdMNjYVf+AotMRNWhW6dyBwAzYI245AXBIO1zB8Pduq2DwCvwvf/Av/dW0ZhH2zfJOyZFzCTRaslDJMFEzspWHENZ85PSVGZHlPxnOHJFNDpeXRLMR2qM1Enxr8Z14TCauu0rUtOi7GA0WsQ8zs4ldlZNXO1Sv1NXMO3uSWbnUeY0HIT6EqVW0ipAFbr2jbZHLzuH4AcLg==; _gid=GA1.2.466490370.1640827957; MEIQIA_VISIT_ID=22z6EbC1k8505p6siaf4v0ivreZ",
	}
	body, result, err := tools.Request(url, method, data, headers, headers)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("状态码：", result.StatusCode)
	fmt.Println(err)
	fmt.Println(string(body))
	fmt.Println(result.Header)
}
