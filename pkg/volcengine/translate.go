package volcengine

import (
	"Translate/tools"
	"net/http"
	"net/url"
	"strings"
)

var host = "https://translate.volcengine.com"
var userAgent = tools.UserAgent()

func Handle(params map[string]string) {
	client := &http.Client{}
	//post要提交的数据
	DataUrlVal := url.Values{}
	DataUrlVal.Add("msToken", "")
	DataUrlVal.Add("X-Bogus", "DFSzswVLQDaokqMPtcHlYuKVuMVl")
	DataUrlVal.Add("_signature", "")

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
	request.Header.Add("Origin", "https://translate.volcengine.com")
	request.Header.Add("Referer", "https://translate.volcengine.com")

	response := tools.Request(*client, *request)
	return response.Body
}
