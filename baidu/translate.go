package baidu

import (
	"fmt"
	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/dop251/goja"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"regexp"
	"strings"
)

const host = "https://fanyi.baidu.com/v2transapi"
const hostDetect = "https://fanyi.baidu.com/langdetect"

var client = initClient()
var html = initHtml()
var jsCompilerVM = goja.New()
var config = initConfig()

func initClient() *http.Client {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}
	return client
}
func initHtml() string {
	request, _ := http.NewRequest("GET", "https://fanyi.baidu.com", nil)
	request.Header.Add("user-agent", browser.Computer())
	//因为首次没有cookie 是不会返回token的，所以请求2次
	_, _ = client.Do(request)
	response, _ := client.Do(request)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
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
	Gtk = "320305.131321201"
	Token = "7b9b8abe67903c95dab61e444ea6e34e"
	config := Config{Token, Gtk}
	return &config
}
func generateSign(query string, config Config) string {
	jsWindow := make(map[string]string)
	jsWindow["gtk"] = config.Gtk
	err := jsCompilerVM.Set("window", jsWindow)
	jsFile := "D:\\goWorkspace\\translate\\baidu\\lib\\sign.js"
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
	request.Header.Add("user-agent", browser.Computer())
	query := request.URL.Query()
	query.Add("query", params["query"])
	request.URL.RawQuery = query.Encode()
	response, _ := client.Do(request)

	detect := FormatDetect(*response)
	return detect
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

	request, err := http.NewRequest("POST", host, strings.NewReader(DataUrlVal.Encode()))
	fmt.Println(DataUrlVal.Encode())
	if err != nil {
		panic(err)
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Host", "fanyi.baidu.com")
	request.Header.Add("Acs-Token", "1663743825590_1663768968232_tUQEsWd7R+vKqFiQIq1ZW8rav2vgE09FBnQG61Mc28otayg4Ov306i5GvlnhOerDwoqkxeTcHGNq9WRkdfI0Y2mkA1pG8/YRTtysB43WKlHBHQs2Qv3N+2ywTL7uUb02NH6tUPDc9KN8WwMal86w2AiGaq4ILlNnqLxshS647lqET1okQ6Egytw5oke9d5UTabRVjxwuQ888wF0PKqyD2nIsbDeq86qDHDMAKEZFKueSRB/yG5RUDkL+klupusnI13PgOuloGEAA5+LMgHZMhjwsGKQ1auGi1slo//aMfltahgAjxbgPrjyNJVrBtxQ+")
	request.Header.Add("Origin", "https://fanyi.baidu.com")
	request.Header.Add("Referer", "https://fanyi.baidu.com/")
	request.Header.Add("User-Agent", browser.Computer())
	request.Header.Add("X-Requested-With", "XMLHttpRequest")

	query := request.URL.Query()
	query.Add("from", source)
	query.Add("to", params["target"])
	request.URL.RawQuery = query.Encode()
	fmt.Println(query.Encode())
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
