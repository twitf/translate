package alibaba

import (
	"Translate/tools"
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
)

var host = "https://translate.alibaba.com/api/translate/text"
var csrfHost = "https://translate.alibaba.com/api/translate/csrftoken"
var userAgent = tools.UserAgent()

func getCsrfToken() Csrf {
	response, err := http.Get(csrfHost)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)
	var csrf Csrf
	tools.FormatResponse(*response, &csrf)
	return csrf
}

func Handle(params map[string]string) Result {
	var csrfToken = getCsrfToken()

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("srcLang", params["source"])
	_ = writer.WriteField("tgtLang", params["target"])
	_ = writer.WriteField("domain", "general")
	_ = writer.WriteField("query", params["query"])
	_ = writer.WriteField(csrfToken.ParameterName, csrfToken.Token)
	err := writer.Close()
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", host, payload)
	if err != nil {
		panic(err)
	}

	request.Header.Add("user-agent", userAgent)
	request.Header.Add(csrfToken.HeaderName, csrfToken.Token)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	response := tools.Request(http.Client{}, *request)
	var result Result
	tools.FormatResponse(response, &result)
	return result
}
