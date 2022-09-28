package tencent

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Result struct {
	SessionUUID string `json:"sessionUuid"`
	Translate   struct {
		ErrCode     int    `json:"errCode"`
		ErrMsg      string `json:"errMsg"`
		SessionUUID string `json:"sessionUuid"`
		Source      string `json:"source"`
		Target      string `json:"target"`
		Records     []struct {
			SourceText string `json:"sourceText"`
			TargetText string `json:"targetText"`
			TraceID    string `json:"traceId"`
		} `json:"records"`
		Full    bool `json:"full"`
		Options struct {
		} `json:"options"`
	} `json:"translate"`
	Dict struct {
		Data []struct {
			Word   string `json:"word"`
			EnHash string `json:"en_hash,omitempty"`
		} `json:"data"`
		ErrCode int    `json:"errCode"`
		ErrMsg  string `json:"errMsg"`
		Type    string `json:"type"`
		Map     struct {
			Life struct {
			} `json:"Life"`
			Is struct {
			} `json:"is"`
			Not struct {
			} `json:"not"`
			Only struct {
			} `json:"only"`
			The struct {
				DetailID string `json:"detailId"`
			} `json:"the"`
			Present struct {
			} `json:"present"`
			But struct {
			} `json:"but"`
			Also struct {
			} `json:"also"`
			Tomorrow struct {
				DetailID string `json:"detailId"`
			} `json:"tomorrow"`
			And struct {
				DetailID string `json:"detailId"`
			} `json:"and"`
			Day struct {
				DetailID string `json:"detailId"`
			} `json:"day"`
			After struct {
				DetailID string `json:"detailId"`
			} `json:"after"`
		} `json:"map"`
	} `json:"dict"`
	Suggest interface{} `json:"suggest"`
	ErrCode int         `json:"errCode"`
	ErrMsg  string      `json:"errMsg"`
}

type Config struct {
	Qtv string `json:"qtv"`
	Qtk string `json:"qtk"`
}

func FormatConfig(response http.Response) Config {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	var config Config
	// 反序列化
	err = json.Unmarshal(body, &config)
	if err != nil {
		fmt.Println(err)
	}
	return config
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
