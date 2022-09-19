package bing

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Result []struct {
	DetectedLanguage struct {
		Language string  `json:"language"`
		Score    float64 `json:"score"`
	} `json:"detectedLanguage"`
	Translations []struct {
		Text            string `json:"text"`
		Transliteration struct {
			Text   string `json:"text"`
			Script string `json:"script"`
		} `json:"transliteration"`
		To      string `json:"to"`
		SentLen struct {
			SrcSentLen   []int `json:"srcSentLen"`
			TransSentLen []int `json:"transSentLen"`
		} `json:"sentLen"`
	} `json:"translations"`
}

type Config struct {
	Key   string
	Token string
	IG    string
	IID   string
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
