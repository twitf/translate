package google

import (
	"Translate/utils"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const host = "https://translate.google.cn/_/TranslateWebserverUi/data/batchexecute"

var userAgent = utils.UserAgent()
var client = initClient()

func Handle(params map[string]string) string {
	config := getConfig()

	requestBody := `[[["MkEWBc","[[\"` + params["query"] + `\",\"` + "auto" + `\",\"` + params["target"] + `\",true],[null]]",null,"generic"]]]`

	DataUrlVal := url.Values{}
	DataUrlVal.Add("f.req", requestBody)

	request, err := http.NewRequest("POST", host, strings.NewReader(DataUrlVal.Encode()))

	if err != nil {
		fmt.Println(err)
	}

	query := request.URL.Query()
	query.Add("rpcids", "MkEWBc")
	query.Add("source-path", "/")
	query.Add("f.sid", config.Fsid)
	query.Add("bl", config.Bl)
	query.Add("hl", params["source"])
	query.Add("soc-app", "1")
	query.Add("soc-platform", "1")
	query.Add("soc-device", "1")
	query.Add("_reqid", getReqId())
	query.Add("rt", "c")
	request.URL.RawQuery = query.Encode()

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	request.Header.Add("User-Agent", userAgent)
	request.Header.Add("host", "translate.google.cn")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func getReqId() string {
	date := time.Now()
	et := 3600*date.Hour() + 60*date.Minute() + date.Second()
	return strconv.Itoa(1 + et + 1e5*1)
}
func initClient() *http.Client {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}
	return client
}

func getConfig() *Config {
	request, _ := http.NewRequest("GET", "https://translate.google.cn/", nil)
	request.Header.Add("user-agent", userAgent)
	response, _ := client.Do(request)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	html := string(body)

	var reg = regexp.MustCompile(`FdrFJe":"(.*?)"`)
	match := reg.FindStringSubmatch(html)
	fsid := match[1]

	var reg2 = regexp.MustCompile(`cfb2h":"(.*?)"`)
	match = reg2.FindStringSubmatch(html)
	//IID 后面的数字是翻译次数 当前是单次翻译固定为1即可
	bl := match[1]

	var config = Config{fsid, bl}
	return &config
}
