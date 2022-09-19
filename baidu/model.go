package baidu

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Csrf struct {
	Token         string `json:"token"`
	ParameterName string `json:"parameterName"`
	HeaderName    string `json:"headerName"`
}

type Result struct {
	RequestID      string `json:"requestId"`
	Success        bool   `json:"success"`
	HTTPStatusCode int    `json:"httpStatusCode"`
	Code           string `json:"code"`
	Message        string `json:"message"`
	Data           struct {
		TranslateText  string `json:"translateText"`
		DetectLanguage string `json:"detectLanguage"`
	} `json:"data"`
}

func FormatResult(response http.Response) Result {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	var result Result
	// 反序列化
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func FormatCsrf(response http.Response) Csrf {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	var csrf Csrf
	// 反序列化
	err = json.Unmarshal(body, &csrf)
	if err != nil {
		fmt.Println(err)
	}
	return csrf
}
