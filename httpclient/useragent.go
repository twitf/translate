package httpclient

import (
	"github.com/DataHenHQ/useragent"
)

func UserAgent() string {
	userAgent, err := useragent.Desktop()
	if err != nil {
		panic(err)
	}
	return userAgent
}
