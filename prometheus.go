package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 定义初始全局变量
var (
	nameSpace string
	subsystem string
)

// 定义用于存储所有指标的 Map
type Metrics map[string][]metriceInfo

// 定义指标内容结构体
type metriceInfo struct {
	MetricersName string
	Help          string
	Type          string
	Labels        map[string]string
	Value         float64
}

// 指定基本信息，并初始化客户端
func InitPrometheus(NameSpace string, SubSystem string) Metrics {
	nameSpace = NameSpace
	subsystem = SubSystem

	metricser := make(Metrics)
	return metricser
}

// 添加 Gauge 指标（静态添加）
// func AddGaugeVec(Name string, Help string, Labels map[string]string, Value float64) {
func (metricser Metrics) AddGaugeVec(Name string, Help string, Labels map[string]string, Value float64) {
	MetricersName := nameSpace + "_" + subsystem + "_" + Name

	MetricersName = strings.Replace(MetricersName, "__", "_", 1)

	metriceInfoer := metriceInfo{
		MetricersName: MetricersName,
		Help:          Help,
		Type:          "gauge",
		Labels:        Labels,
		Value:         Value,
	}

	metricser[MetricersName] = append(metricser[MetricersName], metriceInfoer)
}

// 将标签转换为字符串，以便后续使用
func labels2String(param map[string]string) string {
	var labelsSub []string
	for k, v := range param {
		labelsSub = append(labelsSub, fmt.Sprintf(`%v="%v"`, k, v))
	}
	return (strings.Join(labelsSub, ","))
}

// 将所有指标拼接为多行字符串输出
func (metricser Metrics) getMetrics() string {
	var Metrics string
	for MetricersName, v := range metricser {
		Metrics = Metrics + fmt.Sprintf("# HELP %v %v\n", MetricersName, v[0].Help)
		Metrics = Metrics + fmt.Sprintf("# TYPE %v %v\n", MetricersName, v[0].Type)
		for _, vv := range v {
			Metrics = Metrics + fmt.Sprintf("%v{%v} %v\n", MetricersName, labels2String(vv.Labels), vv.Value)
		}
	}
	return Metrics
}

// 直接打印所有指标
func (metrics Metrics) PrintMetrics() {
	fmt.Println(metrics.getMetrics())
}

// 将结果编码为 Prometheus 文本格式
func (metricser Metrics) WriteTextfile(filename string) error {
	if len(filename) == 0 {
		filename = filepath.Join("/usr/local/node_exporter/textfile_collector/", nameSpace+"-"+subsystem+".prom")
	}
	log.Println("Prom 文件路径：", filename)
	tmp, err := ioutil.TempFile(filepath.Dir(filename), filepath.Base(filename))
	if err != nil {
		return err
	}

	labels := map[string]string{
		"NameSpace": nameSpace,
	}

	metricser.AddGaugeVec("script_last_update_time", "脚本最后更新时间", labels, float64(time.Now().Unix()))

	// 将内容写入文件
	if _, err := tmp.Write([]byte(metricser.getMetrics())); err != nil {
		fmt.Println(err)
	}

	// 延迟关闭文件
	defer os.Remove(tmp.Name())

	// 关闭文件
	if err := tmp.Close(); err != nil {
		return err
	}

	// 修改文件权限
	if err := os.Chmod(tmp.Name(), 0644); err != nil {
		return err
	}
	// 返回文件名
	return os.Rename(tmp.Name(), filename)
}
