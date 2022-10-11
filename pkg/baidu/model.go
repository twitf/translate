package baidu

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Config struct {
	Token string
	Gtk   string
}
type Detect struct {
	Error int    `json:"error"`
	Msg   string `json:"msg"`
	Lan   string `json:"lan"`
}

type Result struct {
	TransResult struct {
		Data []struct {
			Dst        string          `json:"dst"`
			PrefixWrap int             `json:"prefixWrap"`
			Result     [][]interface{} `json:"result"`
			Src        string          `json:"src"`
		} `json:"data"`
		From   string `json:"from"`
		Status int    `json:"status"`
		To     string `json:"to"`
		Type   int    `json:"type"`
	} `json:"trans_result"`
	LijuResult struct {
		Double string `json:"double"`
		Single string `json:"single"`
	} `json:"liju_result"`
	Logid int `json:"logid"`
}

func FormatDetect(response http.Response) Detect {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	var detect Detect
	// 反序列化
	err = json.Unmarshal(body, &detect)
	if err != nil {
		fmt.Println(err)
	}
	return detect
}
