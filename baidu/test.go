package baidu

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

var cookies_lagou []*http.Cookie

const (
	login_url_lagou string = "https://passport.lagou.com/login/login.html"

	post_login_info_url_lagou string = "https://passport.lagou.com/login/login.json"

	username_lagou string = "13330295142"
	password_lagou string = "4525674692ac06e619cdb3f1b4b65b08"
)

func getToken1(contents io.Reader) (string, string) {

	data, _ := ioutil.ReadAll(contents)

	regCode := regexp.MustCompile(`X_Anti_Forge_Code = '(.*?)';`)
	if regCode == nil {
		log.Fatal("解析Code出错...")
	}

	//提取关键信息
	code := regCode.FindAllStringSubmatch(string(data), -1)[0][1]

	regToken := regexp.MustCompile(`X_Anti_Forge_Token = '(.*?)';`)
	if regToken == nil {
		fmt.Println("MustCompile err")
	}

	//提取关键信息
	token := regToken.FindAllStringSubmatch(string(data), -1)[0][1]

	return token, code
}

func login_lagou() {
	//获取登陆界面的cookie
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}
	req, _ := http.NewRequest("GET", login_url_lagou, nil)
	res, _ := client.Do(req)

	token, code := getToken1(res.Body)
	fmt.Println(token, code)
	//post数据
	postValues := url.Values{}
	postValues.Add("isValidate", "true")
	postValues.Add("username", username_lagou)
	postValues.Add("password", password_lagou)
	postValues.Add("request_form_verifyCode", "")
	postValues.Add("submit", "")
	// body := ioutil.NopCloser(strings.NewReader(postValues.Encode())) //把form数据编下码
	// requ, _ := http.NewRequest("POST", post_login_info_url_lagou, nil)

	requ, _ := http.NewRequest("POST", post_login_info_url_lagou, strings.NewReader(postValues.Encode()))
	requ.Header.Set("Referer", "https://passport.lagou.com/login/login.html")
	requ.Header.Set("X-Requested-With", "XMLHttpRequest")
	requ.Header.Set("X-Anit-Forge-Token", token)
	requ.Header.Set("X-Anit-Forge-Code", code)
	requ.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:51.0) Gecko/20100101 Firefox/51.0")
	requ.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	//for _, v := range res.Cookies() {
	//    requ.AddCookie(v)
	//}

	res, _ = client.Do(requ)
	//cookies_lagou = res.Cookies()
	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println(string(data))
}

func main() {
	login_lagou()
}
