package youdao

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	browser "github.com/EDDYCJY/fake-useragent"
	"io"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const host = "https://fanyi.youdao.com/translate"

var userAgent = browser.Computer()
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
	response, _ := client.Do(request)
	fmt.Println(response.Header)
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
	sign := Md5("fanyideskweb" + salt + query + "Ygy_4c=r#e#4EX^NUGUc5")
	return Config{Bv: bv, Lts: lts, Salt: salt, Sign: sign}

}

func Handle(params map[string]string) Result {
	initCookie()
	config := generateConfig(params["query"])
	fmt.Println(config)
	//post要提交的数据
	DataUrlVal := url.Values{}
	DataUrlVal.Add("i", params["query"])
	DataUrlVal.Add("from", params["source"])
	DataUrlVal.Add("to", params["target"])
	DataUrlVal.Add("smartresult", "dict")
	DataUrlVal.Add("client", "fanyideskweb")
	DataUrlVal.Add("salt", config.Salt)
	DataUrlVal.Add("sign", config.Sign)
	DataUrlVal.Add("lts", config.Lts)
	DataUrlVal.Add("bv", config.Bv)
	DataUrlVal.Add("doctype", "json")
	DataUrlVal.Add("version", "2.1")
	DataUrlVal.Add("keyfrom", "fanyi.web")
	DataUrlVal.Add("action", "FY_BY_REALTlME")

	request, err := http.NewRequest("POST", host, strings.NewReader(DataUrlVal.Encode()))
	if err != nil {
		panic(err)
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Add("Host", "fanyi.youdao.com")
	request.Header.Add("User-Agent", userAgent)
	request.Header.Add("Referer", "https://fanyi.youdao.com/")
	request.Header.Add("Cookie", "OUTFOX_SEARCH_USER_ID_NCOO=1519430159.539675; OUTFOX_SEARCH_USER_ID=\"1233087801@10.108.162.133\"; ___rl__test__cookies=1663837506878")

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
