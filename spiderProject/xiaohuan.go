package spiderProject

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"proxy/pkg/htmlquery"
	"strings"
)


func ihuanWork() {
	client := &http.Client{}
	var data = strings.NewReader(`num=100&port=&kill_port=&address=%E4%B8%AD%E5%9B%BD&kill_address=&anonymity=&type=0&post=1&sort=1&key=174952c9c4e127d310bc4b298f04645f`)
	req, err := http.NewRequest("POST", "https://ip.ihuan.me/tqdl.html", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "ip.ihuan.me")
	req.Header.Set("cache-control", "max-age=0")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36")
	req.Header.Set("origin", "https://ip.ihuan.me")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("referer", "https://ip.ihuan.me/ti.html")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("cookie", "__cfduid=d54a595894736d3f847007bfafe8cd7d71597036698; statistics=50f6bea7eecdc0cbebce3fe52b001651")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc := htmlquery.ReadFromString(string(bodyText))
	list := htmlquery.QueryAll(doc, "/html/body/meta\"utf-8\"/div[2]/div/div[2]")

}
