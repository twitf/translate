package baidu

import (
	"Translate/tools"
	"fmt"
	"github.com/dop251/goja"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const host = "https://fanyi.baidu.com/v2transapi"
const hostDetect = "https://fanyi.baidu.com/langdetect"

var userAgent = tools.UserAgent()
var client = tools.Client()
var html = initHtml()
var jsCompilerVM = goja.New()
var config = initConfig()

func initHtml() string {
	request, _ := http.NewRequest("GET", "https://fanyi.baidu.com", nil)
	request.Header.Add("user-agent", userAgent)
	//因为首次没有cookie 是不会返回token的，所以请求2次
	_, _ = client.Do(request)
	response, _ := client.Do(request)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	html := string(body)
	return html
}

func initConfig() *Config {
	regGtk := regexp.MustCompile(`window.gtk = "(.*?)";`)
	matchGtk := regGtk.FindStringSubmatch(html)
	Gtk := matchGtk[1]

	regToken := regexp.MustCompile(`token: '(.*?)',`)
	matchToken := regToken.FindStringSubmatch(html)
	Token := matchToken[1]

	config := Config{Token, Gtk}
	return &config
}
func generateSign(query string, config Config) string {
	jsWindow := make(map[string]string)
	jsWindow["gtk"] = config.Gtk
	err := jsCompilerVM.Set("window", jsWindow)

	path, _ := os.Getwd()
	jsFile := path + "/pkg/baidu/lib/sign.js"
	bytes, err := os.ReadFile(jsFile)
	if err != nil {
		panic(err)
	}

	_, err = jsCompilerVM.RunString(string(bytes))
	getSign, ok := goja.AssertFunction(jsCompilerVM.Get("getSign"))
	if !ok {
		panic("Not a function getSign")
	}
	sign, err := getSign(goja.Undefined(), jsCompilerVM.ToValue(query))
	if err != nil {
		panic(err)
	}
	return sign.ToString().String()
}
func getDetect(params map[string]string) Detect {
	request, _ := http.NewRequest("POST", hostDetect, nil)
	request.Header.Add("user-agent", userAgent)
	query := request.URL.Query()
	query.Add("query", params["query"])
	request.URL.RawQuery = query.Encode()
	response, _ := client.Do(request)

	var detect Detect
	tools.FormatResponse(*response, &detect)
	return detect
}

func getTimestamp() string {
	return strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
}

func Handle(params map[string]string) Result {
	var source string
	source, ok := params["source"]
	if ok == false {
		detect := getDetect(params)
		source = detect.Lan
	}

	sign := generateSign(params["query"], *config)

	//post要提交的数据
	DataUrlVal := url.Values{}
	DataUrlVal.Add("from", source)
	DataUrlVal.Add("to", params["target"])
	DataUrlVal.Add("query", params["query"])
	DataUrlVal.Add("transtype", "realtime")
	DataUrlVal.Add("simple_means_flag", "3")
	DataUrlVal.Add("sign", sign)
	DataUrlVal.Add("token", config.Token)
	DataUrlVal.Add("domain", "common")
	DataUrlVal.Add("ts", getTimestamp())

	request, err := http.NewRequest("POST", host, strings.NewReader(DataUrlVal.Encode()))
	if err != nil {
		panic(err)
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Add("Host", "fanyi.baidu.com")
	request.Header.Add("User-Agent", userAgent)

	query := request.URL.Query()
	query.Add("from", source)
	query.Add("to", params["target"])
	request.URL.RawQuery = query.Encode()

	response := tools.Request(*client, *request)
	fmt.Println(response)
	var result Result
	tools.FormatResponse(response, &result)
	return result
}
