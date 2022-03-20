package spiderProject

import (
	"io/ioutil"
	"log"
	"net/http"
	"proxy/internal/proxy"
	"proxy/pkg/htmlquery"
)

func Spider_yun() (infoList []proxy.IPInfo) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://ip.jiangxianli.com/?page=1&country=%E4%B8%AD%E5%9B%BD", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.74 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://ip.jiangxianli.com/?page=1&country=%E4%B8%AD%E5%9B%BD")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc := htmlquery.ReadFromString(string(bodyText))
	tr := htmlquery.QueryAll(doc, "/html/body/div[1]/div[2]/div[1]/div[1]/table/tbody/tr")
	for _, n := range tr {

		if len(htmlquery.FindOneText(n, "//td[1]")) <= 15 {
			var info = proxy.IPInfo{
				IP:     htmlquery.FindOneText(n, "//td[1]"),
				Port:   htmlquery.FindOneText(n, "//td[2]"),
				IPType: switchIPType(htmlquery.FindOneText(n, "//td[4]")),
				Rating: 50,
			}
			infoList = append(infoList, info)
		}

	}
	return infoList
}
