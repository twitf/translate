package baidu

import (
	"bytes"
	"github.com/corpix/uarand"
	"io"
	"mime/multipart"
	"net/http"
)

var host = "https://fanyi.baidu.com/v2transapi"

func Handle(params map[string]string) Result {

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("srcLang", params["source"])
	_ = writer.WriteField("tgtLang", params["target"])
	_ = writer.WriteField("domain", "general")
	_ = writer.WriteField("query", params["query"])
	err := writer.Close()
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", host, payload)
	if err != nil {
		panic(err)
	}

	request.Header.Add("user-agent", uarand.GetRandom())
	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
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
