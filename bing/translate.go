package bing

import (
	"encoding/json"
	browser "github.com/EDDYCJY/fake-useragent"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var host = "https://cn.bing.com/ttranslatev3"
var userAgent = browser.Computer()

func Handle(params map[string]string) Result {
	config := getConfig()

	client := &http.Client{}
	//post要提交的数据
	DataUrlVal := url.Values{}
	DataUrlVal.Add("fromLang", params["source"])
	DataUrlVal.Add("text", params["query"])
	DataUrlVal.Add("to", params["target"])
	DataUrlVal.Add("token", config.Token)
	DataUrlVal.Add("key", config.Key)

	request, err := http.NewRequest("POST", host, strings.NewReader(DataUrlVal.Encode()))

	if err != nil {
		panic(err)
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("user-agent", userAgent)

	query := request.URL.Query()
	query.Add("isVertical", "1")
	query.Add("IG", config.IG)
	query.Add("IID", config.IID)
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

func getConfig() *Config {

	res, err := http.Get("https://cn.bing.com/translator")
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	html := string(body)
	re := regexp.MustCompile(`var params_RichTranslateHelper = (.*?);`)
	match := re.FindStringSubmatch(html)

	strArr := make([]interface{}, 11)
	err = json.Unmarshal([]byte(match[1]), &strArr)
	if err != nil {
		panic(err)
	}
	//断言，顾名思义就是果断的去猜测一个未知的事物。在 go 语言中，interface{} 就是这个神秘的未知类型，其断言操作就是用来判断 interface{} 的类型。
	floatKey := strArr[0].(float64)
	key := strconv.FormatFloat(floatKey, 'f', 0, 64)
	token := strArr[1].(string)

	var reg = regexp.MustCompile(`IG:"(.*?)"`)
	match = reg.FindStringSubmatch(html)
	IG := match[1]

	var reg2 = regexp.MustCompile(`<div id="rich_tta" data-iid="(.*?)"`)
	match = reg2.FindStringSubmatch(html)
	IID := match[1]

	var config = Config{key, token, IG, IID}
	return &config
}
