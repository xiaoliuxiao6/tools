package tools_test

import (
	"testing"

	"github.com/xiaoliuxiao6/tools/v2"
)

// 使用方法演示
func TestUsege(t *testing.T) {
	// 1.注册指标
	metrics := tools.InitPrometheus("opensea", "gala")

	// 2.添加指标
	labels := map[string]string{
		"type": "1",
	}
	metrics.AddGaugeVec("测试1", "帮助内容1", labels, 1)
	metrics.AddGaugeVec("测试11", "帮助内容11", labels, 1)

	// 3.（可选）打印指标
	metrics.PrintMetrics()

	// 3.（可选）将指标写入文件
	metrics.WriteTextfile("")
}
