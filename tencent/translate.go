package tencent

import (
	"Translate/utils"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var host = "https://fanyi.qq.com/api/translate"
var userAgent = utils.UserAgent()

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

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)
	return FormatResult(*response)
}

func getConfig() *Config {
	response, err := http.Post("https://fanyi.qq.com/api/reauth12f", "", nil)
	if err != nil {
		panic(err)
	}
	config := FormatConfig(*response)
	return &config
}
