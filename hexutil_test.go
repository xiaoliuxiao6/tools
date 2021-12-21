package tools_test

import (
	"fmt"
	"testing"

	"github.com/xiaoliuxiao6/tools"
)

// 将 0x 开头的16进制字符串解析为 unint64
func TestDecodeUint64(t *testing.T) {
	unint64Sting := "0xd27b8c"
	unint64, err := tools.DecodeUint64(unint64Sting)
	if err != nil {
		fmt.Println("解析错误")
	}
	fmt.Printf("字符串为：%v, 解析后结果为：%v\n", unint64Sting, unint64)
}
