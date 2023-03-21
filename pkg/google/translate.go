package google

import (
	"Translate/tools"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const host = "https://translate.google.com/_/TranslateWebserverUi/data/batchexecute"

var userAgent = tools.UserAgent()
var client = tools.Client()

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
	request.Header.Add("X-Goog-BatchExecute-Bgr", "[\";FQu4C0bQAAYnbWexB21fFeuymVt2YhgmACkAIwj8RsdkWB8o17-W5-2JIXa6FtZCU1974Z_feyjAz0tQmsmI6ZSahR8AAABITwAAAAF1AQeEAqw_uOkDvKtmFWcReh8Yw03zweEFu5khuV-3EEH3wqleSkRq_kIYrVKaq6MZ5oQmyQAPgUfBDoy5m0A0ofICKzgbfMzovzyVj6tiwcCs2ooJSJtWEJ0Y_2keoba8AZU0TzcDp_DJQbN1WOM0iObq3YqY8og-RzW2tCJv7LoBa1zIj1aw9B3IKFJyjooELCL7Ynpgp1MDxbultsRZegkvidvIDxYmv9OweX1EnWsF3SvSPuteiXLaTtl3k_iCgsZTGRWhpX-U47XFcK7dB5cvTQufALIPsgMc9sB9Cn8e_So8jtD5c489l2tExHUg6xpMdByYYyFyVuinfxsuCwO64y9chgroz8mt34meYGtj-VBBYq8ReI5YV_LkfMF2zxnX08frGilyUnDr1-QpzsaOQQ0hbbA3Cewkoe2huuv-sgVzNMvkfZ0m72XNFTMZeO5ZhdgqPU2RtJWx8ZwP_VJNhnHX74W4CNdEe0gMLGPbI_rtWOnB3VOuVlG-d6G_TV_Qx6y2i5EfzatnU3jGVVEReVqvlNFnTULfnFYCShrYVgbvoiLnhnbuwg7u3NJ-VqdUFTAoAQdJcMO8VJbM1fDTXpX_R4MYcBbO8NcRqkwr_SmvOI0kkkCqqWH1jW3YW-9ba5AD--lPy3tV-nhaPWT4CGrW5C037qUTkxX4YrnwtI3rFOJSOVO-Yj2FgaUgK3drVxYz2kFK6Qx92XJHi2KBjtuCXPQn5Gg4duqotw3uBQbyCVqX3KkEJi6ws3u5iVVnLKMrWR03mUlQSPy7vA4spbZZ7Gk4Bu4L-0lb7Mtp0mrMOdk8_G1M_D2wvkoS0UzxyuwEf4XBHT1m0nQjin9p9q1Cgyhatw0K___w-qEqImKqWfSLSFzuSFkf8R-7p7ZGy6dhZL5dZkuPoGdTtF8\",null,null,11,null,null,null,0,\"2\"]")
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
	return getTranslation(string(body))
}

func getReqId() string {
	date := time.Now()
	et := 3600*date.Hour() + 60*date.Minute() + date.Second()
	return strconv.Itoa(1 + et + 1e5*1)
}

func getConfig() *Config {
	request, _ := http.NewRequest("GET", "https://translate.google.com/", nil)
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
	bl := match[1]

	var config = Config{fsid, bl}
	return &config
}

func getTranslation(body string) string {
	data := strings.Split(body, "\n")
	str := strings.ReplaceAll(data[3], "\\", "")
	str = strings.ReplaceAll(str, "\"[", "[")
	str = strings.ReplaceAll(str, "]\"", "]")

	value := gjson.Get(str, "0.2.1.0.0.5.0.0")
	return value.String()
}
