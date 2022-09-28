package youdao

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Config struct {
	Bv   string
	Lts  string
	Salt string
	Sign string
}

type Result struct {
	ErrorCode       int `json:"errorCode"`
	TranslateResult [][]struct {
		Tgt string `json:"tgt"`
		Src string `json:"src"`
	} `json:"translateResult"`
	Type        string `json:"type"`
	SmartResult struct {
		Entries []string `json:"entries"`
		Type    int      `json:"type"`
	} `json:"smartResult"`
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
