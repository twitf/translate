package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	url := "https://translate.google.com/"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("sec-ch-ua", "\"Chromium\";v=\"106\", \"Google Chrome\";v=\"106\", \"Not;A=Brand\";v=\"99\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-full-version", "\"106.0.5249.103\"")
	req.Header.Add("sec-ch-ua-arch", "\"x86\"")
	req.Header.Add("sec-ch-ua-platform", "\"Linux\"")
	req.Header.Add("sec-ch-ua-platform-version", "\"5.15.45\"")
	req.Header.Add("sec-ch-ua-model", "\"\"")
	req.Header.Add("sec-ch-ua-bitness", "\"64\"")
	req.Header.Add("sec-ch-ua-wow64", "?0")
	req.Header.Add("sec-ch-ua-full-version-list", "\"Chromium\";v=\"106.0.5249.103\", \"Google Chrome\";v=\"106.0.5249.103\", \"Not;A=Brand\";v=\"99.0.0.0\"")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("X-Client-Data", "CKS1yQEIk7bJAQimtskBCKmdygEI6enKAQiWocsBCKu8zAEIzbzMAQjLxswBCOLLzAEIht3MAQif38wBCPLfzAEI8ODMAQjD4cwBCMXhzAEIyePMAQic5MwBCMfmzAE=")
	req.Header.Add("Sec-Fetch-Site", "none")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("host", "translate.google.com")
	req.Header.Add("Cookie", "NID=511=I3ehRZVQIVyCjgDaD7yTFHAWwjBEYNSuvb120OiSKXyKft2xBa2cwVzI7CsHd1qmWwEssBkZJ1nmQPL_KMjdntL8684hyrbnR5ABPvX6oSLpnQOOq2QK-dMVLg93PUzXMNL6qSff2IbPa-j-iDl9RUn6VLWgTexRwRBx0kBKxwI")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err2 := ioutil.WriteFile("./output2.html", body, 0666) //写入文件(字节数组)
	if err2 != nil {
		log.Fatal(err2.Error())
	}
}
