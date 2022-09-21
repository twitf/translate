package baidu

import (
	"fmt"
	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/dop251/goja"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
	"regexp"
)

const host = "https://fanyi.baidu.com/v2transapi"

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

	config := Config{Token, Gtk}
	return &config
}
func generateSign(query string, config Config) string {
	jsWindow := make(map[string]string)
	jsWindow["gtk"] = config.Gtk
	err := jsCompilerVM.Set("window", jsWindow)
	jsFile := "/home/administrator/GolandProjects/translate/baidu/lib/sign.js"
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
func getDetect() {

}
func Handle(params map[string]string) {
	sign := generateSign(params["query"], *config)
	fmt.Println(sign)
	//client := &http.Client{}
	////post要提交的数据
	//DataUrlVal := url.Values{}
	//DataUrlVal.Add("fromLang", params["source"])
	//DataUrlVal.Add("text", params["query"])
	//DataUrlVal.Add("to", params["target"])
	////DataUrlVal.Add("token", config.Token)
	////DataUrlVal.Add("key", config.Key)
	//
	//request, err := http.NewRequest("POST", host, strings.NewReader(DataUrlVal.Encode()))
	//
	//if err != nil {
	//	panic(err)
	//}
	//request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//request.Header.Add("user-agent", uarand.GetRandom())
	//
	//query := request.URL.Query()
	//query.Add("isVertical", "1")
	////query.Add("IG", config.IG)
	////query.Add("IID", config.IID)
	//request.URL.RawQuery = query.Encode()
	//
	//response, err := client.Do(request)
	//if err != nil {
	//	panic(err)
	//}
	//defer func(Body io.ReadCloser) {
	//	err := Body.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(response.Body)
	//return FormatResult(*response)
}
