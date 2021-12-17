package tools

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 遍历结构体
func StructFor(input interface{}) {
	inputValue := reflect.ValueOf(input)
	inputType := reflect.TypeOf(input)
	iCount := inputValue.NumField() // 获取结构体有几个字段
	for i := 0; i < iCount; i++ {
		iName := inputType.Field(i).Name           // 字段名称
		iValue := inputValue.Field(i)              // 字段值
		iTag := inputType.Field(i).Tag.Get("json") // 字段标签（string）
		fmt.Printf("字段名称：%v, 字段值：%v, 字段标签：%v\n", iName, iValue, iTag)
	}
}

// 打印结构体
func StructPrint(input interface{}) {
	// 没有格式的打印结构体，同时显示结构体字段和值
	// fmt.Printf("%+v", input)

	// 转换为 JSON 并漂亮输出结构体
	output, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		fmt.Println("发生了错误：", err)
		panic("")
	}
	fmt.Println(string(output))
}
