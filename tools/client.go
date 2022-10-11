package tools

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

func Request(client http.Client, request http.Request) http.Response {
	response, err := client.Do(&request)
	if err != nil {
		panic(err)
	}
	return *response
}
