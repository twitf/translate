#### 目前支持的服务
- [x] 阿里翻译
- [x] 百度翻译
- [x] 必应翻译
- [x] 有道翻译
- [x] 腾讯翻译
- [x] 谷歌翻译(`X-Goog-BatchExecute-Bgr`这个Header影响翻译准确性，尚未完成解析)
- [x] 火山翻译


## 
```go
package main

import (
	"Translate/pkg/alibaba"
	"Translate/pkg/baidu"
	"Translate/pkg/bing"
	"Translate/pkg/google"
	"Translate/pkg/tencent"
	"Translate/pkg/volcengine"
	"Translate/pkg/youdao"
	"fmt"
)

func main() {
	testAlibaba()
	testBing()
	testBaidu()
	testYoudao()
	testTencent()
	testGoogle()
	testVolcengine()
}
func testTencent() {
	params := make(map[string]string)
	params["source"] = "auto"
	params["target"] = "en"
	params["query"] = "生活不止眼前的苟且，还有明天和后天的苟且"
	result := tencent.Handle(params)
	fmt.Println(result.Translate.Records[0].TargetText)
}
func testAlibaba() {
	params := make(map[string]string)
	params["source"] = "auto"
	params["target"] = "en"
	params["query"] = "生活不止眼前的苟且，还有明天和后天的苟且"

	result := alibaba.Handle(params)
	fmt.Println(result.Data.TranslateText)
}
func testBing() {
	params := make(map[string]string)
	params["source"] = "zh-Hans"
	params["target"] = "en"
	params["query"] = "生活不止眼前的苟且，还有明天和后天的苟且"
	result := bing.Handle(params)
	fmt.Println(result[0].Translations[0].Text)
}
func testBaidu() {
	params := make(map[string]string)
	//params["source"] = "auto"
	params["target"] = "en"
	params["query"] = "生活不止眼前的苟且，还有明天和后天的苟且"
	result := baidu.Handle(params)
	fmt.Println(result.TransResult.Data[0].Dst)
}

func testYoudao() {
	params := make(map[string]string)
	params["source"] = "AUTO"
	params["target"] = "AUTO"
	params["query"] = "生活不止眼前的苟且，还有明天和后天的苟且"
	result := youdao.Handle(params)
	fmt.Println(result.TranslateResult[0][0].Tgt)
}

func testGoogle() {
	params := make(map[string]string)
	params["source"] = "zh-CN"
	params["target"] = "en"
	params["query"] = "生活不止眼前的苟且，还有明天和后天的苟且"
	result := google.Handle(params)
	fmt.Println(result)
}

func testVolcengine() {
	params := make(map[string]string)
	params["source"] = "zh"
	params["target"] = "en"
	params["query"] = "生活不止眼前的苟且，还有明天和后天的苟且"
	result := volcengine.Handle(params)
	fmt.Println(result.Translation)
}
```
