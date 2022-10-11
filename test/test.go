package main

import (
	"Translate/pkg/alibaba"
	"Translate/pkg/baidu"
	"Translate/pkg/bing"
	"Translate/pkg/tencent"
	"Translate/pkg/youdao"
	"Translate/tools"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//testAlibaba()
	////国内访问略慢 不建议
	//testBing()
	//testBaidu()
	//testYoudao()
	//testTencent()
	testGoogle()
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
	//params := make(map[string]string)
	//params["source"] = "zh-CN"
	//params["target"] = "en"
	//params["query"] = "生活不止眼前的苟且，还有明天和后天的苟且"
	//result := google.Handle(params)
	result := `)]}'

575
[["wrb.fr","MkEWBc","[[\"Nǐ hǎo\",null,\"zh-CN\",[[[0,[[[null,2]],[true]]]],2],[[\"你好\",null,null,2]]],[[[null,null,null,null,null,[[\"Hello\",null,null,null,[[\"Hello\",[4,5],[]],[\"Hi\",[4,11],[]],[\"Hello there\",[4],[]]]]]]],\"en\",1,\"zh-CN\",[\"你好\",\"auto\",\"en\",true]],\"zh-CN\",[\"你好!\",null,null,null,null,[[[\"感叹词\",[[\"Hello!\",null,[\"你好!\",\"喂!\"],1,true],[\"Hi!\",null,[\"嗨!\",\"你好!\"],1,true],[\"Hallo!\",null,[\"你好!\"],3,true]],\"en\",\"zh-CN\"]],3],null,null,\"zh-CN\",1]]",null,null,null,"generic"],["di",34],["af.httprm",33,"-4556213990339052700",83]]
25
[["e",4,null,null,647]]
`
	data := tools.Explode("\n", result)
	lcc1 := []string(data)
	lcc2 := []string(data)
	fmt.Printf("TYPE is %T，value is %+v\n", lcc1, lcc1)
	fmt.Printf("TYPE is %T，value is %+v\n", lcc2, lcc2)

}

func ReadLine(lineNumber int) string {
	file, _ := os.Open("log.txt")
	strings.SplitN()
	fileScanner := bufio.NewScanner(file)
	lineCount := 1
	for fileScanner.Scan() {
		if lineCount == lineNumber {
			return fileScanner.Text()
		}
		lineCount++
	}
	defer file.Close()
	return ""
}
