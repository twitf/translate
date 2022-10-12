package tools

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FormatResponse(response http.Response, val interface{}) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	JsonDecode(body, &val)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)
}

func JsonDecode(data []byte, val interface{}) {
	err := json.Unmarshal(data, &val)
	if err != nil {
		panic("Json Decode Error:" + err.Error())
	}
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
