package main

import (
	"Translate/alibaba"
	"Translate/baidu"
	"Translate/bing"
	"fmt"
)

func main() {
	testBaidu()
}

func testAlibaba() {
	params := make(map[string]string)
	params["source"] = "auto"
	params["target"] = "en"
	params["query"] = "请注意审核目标那块的驳回继续沿用之前的API，那块和其他的处理逻辑不一样 请加上remark参数"

	result := alibaba.Handle(params)
	fmt.Println(result.Data.TranslateText)
}
func testBing() {
	params := make(map[string]string)
	params["source"] = "zh-Hans"
	params["target"] = "en"
	params["query"] = "请注意审核目标那块的驳回继续沿用之前的API，那块和其他的处理逻辑不一样 请加上remark参数"
	result := bing.Handle(params)
	fmt.Println(result[0].Translations[0].Text)
}
func testBaidu() {
	params := make(map[string]string)
	//params["source"] = "auto"
	params["target"] = "en"
	params["query"] = "你ahao"
	result := baidu.Handle(params)
	fmt.Println(result)
}
