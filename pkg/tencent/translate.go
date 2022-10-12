package tencent

import (
	"Translate/tools"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var host = "https://fanyi.qq.com/api/translate"
var userAgent = tools.UserAgent()

func Handle(params map[string]string) Result {
	config := getConfig()
	client := &http.Client{}
	//post要提交的数据
	DataUrlVal := url.Values{}
	DataUrlVal.Add("qtk", config.Qtk)
	DataUrlVal.Add("qtv", config.Qtv)
	DataUrlVal.Add("randstr", "")
	DataUrlVal.Add("sessionUuid", "translate_uuid"+strconv.FormatInt(time.Now().UnixNano()/1e6, 10))

	DataUrlVal.Add("fromLang", params["source"])
	DataUrlVal.Add("sourceText", params["query"])
	DataUrlVal.Add("target", params["target"])
	DataUrlVal.Add("ticket", "")
	request, err := http.NewRequest("POST", host, strings.NewReader(DataUrlVal.Encode()))

	if err != nil {
		panic(err)
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("user-agent", userAgent)
	request.Header.Add("host", "fanyi.qq.com")
	request.Header.Add("Origin", "https://fanyi.qq.com")
	request.Header.Add("Referer", "https://fanyi.qq.com/")
	request.Header.Add("Cookie", "qtv="+config.Qtv+"; qtk="+config.Qtk)

	response := tools.Request(*client, *request)
	var result Result
	tools.FormatResponse(response, &result)
	return result
}

func getConfig() *Config {
	response, err := http.Post("https://fanyi.qq.com/api/reauth12f", "", nil)
	if err != nil {
		panic(err)
	}
	var config Config
	tools.FormatResponse(*response, &config)
	return &config
}
