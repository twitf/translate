package youdao

import (
	"Translate/utils"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"time"
)

const host = "http://fanyi.youdao.com/translate_o"

var userAgent = utils.UserAgent()
var client = initClient()

func initClient() *http.Client {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}
	return client
}
func initCookie() {
	request, _ := http.NewRequest("GET", "https://fanyi.youdao.com", nil)
	request.Header.Add("user-agent", userAgent)
	_, _ = client.Do(request)
}
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func generateConfig(query string) Config {
	bv := Md5(userAgent)
	lts := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	salt := lts + strconv.Itoa(rand.Intn(10))
	bv = "47edca4d7e6ec9bf4fca7156ea36b8ef"
	sign := Md5("fanyideskweb" + query + salt + "Ygy_4c=r#e#4EX^NUGUc5")
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
