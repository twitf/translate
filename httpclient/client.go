package httpclient

import (
	"net/http"
	"net/http/cookiejar"
)

func Client() *http.Client {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}
	return client
}
