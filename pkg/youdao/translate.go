package youdao

import (
	"Translate/tools"
	"bytes"
	"math/rand"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
)

const host = "http://fanyi.youdao.com/translate_o"

var userAgent = tools.UserAgent()
var client = tools.Client()

func initCookie() {
	request, _ := http.NewRequest("GET", "https://fanyi.youdao.com", nil)
	request.Header.Add("user-agent", userAgent)
	_, _ = client.Do(request)
}

func generateConfig(query string) Config {
	bv := tools.Md5(userAgent)
	lts := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	salt := lts + strconv.Itoa(rand.Intn(10))
	sign := tools.Md5("fanyideskweb" + query + salt + "Ygy_4c=r#e#4EX^NUGUc5")
	return Config{Bv: bv, Lts: lts, Salt: salt, Sign: sign}

}

func Handle(params map[string]string) Result {
	initCookie()
	config := generateConfig(params["query"])
	//post要提交的数据

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("i", params["query"])
	_ = writer.WriteField("from", params["source"])
	_ = writer.WriteField("to", params["target"])
	_ = writer.WriteField("smartresult", "dict")
	_ = writer.WriteField("client", "fanyideskweb")
	_ = writer.WriteField("salt", config.Salt)
	_ = writer.WriteField("sign", config.Sign)
	_ = writer.WriteField("lts", config.Lts)
	_ = writer.WriteField("bv", config.Bv)
	_ = writer.WriteField("doctype", "json")
	_ = writer.WriteField("version", "2.1")
	_ = writer.WriteField("keyfrom", "fanyi.web")
	_ = writer.WriteField("action", "FY_BY_REALTlME")
	err := writer.Close()
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", host, payload)
	if err != nil {
		panic(err)
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.2; rv:51.0) Gecko/20100101 Firefox/51.0")
	request.Header.Add("Referer", "http://fanyi.youdao.com/")
	//request.Header.Add("Cookie", "OUTFOX_SEARCH_USER_ID=-2022895048@10.168.8.76;")
	request.Header.Set("Content-Type", writer.FormDataContentType())

	query := request.URL.Query()
	query.Add("smartresult", "dict")
	query.Add("smartresult", "rule")
	request.URL.RawQuery = query.Encode()

	response := tools.Request(*client, *request)
	var result Result
	tools.FormatResponse(response, &result)
	return result
}
