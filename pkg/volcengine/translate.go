package volcengine

import (
	"Translate/tools"
	"bytes"
	"encoding/json"
	"net/http"
)

var host = "https://translate.volcengine.com/crx/translate/v1"
var userAgent = tools.UserAgent()

func Handle(params map[string]string) Result {
	client := &http.Client{}
	//post要提交的数据
	var requestParam RequestParam
	requestParam.SourceLanguage = params["source"]
	requestParam.TargetLanguage = params["target"]
	requestParam.Text = params["query"]

	jsonBody, err := json.Marshal(requestParam)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", host, bytes.NewBuffer(jsonBody))
	if err != nil {
		panic(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("user-agent", userAgent)
	response := tools.Request(*client, *request)

	var result Result
	tools.FormatResponse(response, &result)
	return result
}
