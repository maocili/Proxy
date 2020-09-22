package spiderProject

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"proxy/internal/proxy"
	"proxy/pkg/htmlquery"
	"time"
)

func switchIPType(t string) int {
	switch t {
	case "HTTP":
		return proxy.HTTP
	case "HTTPS":
		return proxy.HTTPS
	default:
		return 0
	}
}

func MyRecover()  {
	if r := recover(); r != nil {
		fmt.Errorf("Unknown panic: %v", r)
	}
}

func Spider_xl()(infoList []proxy.IPInfo ){
	log.Println("Running Spider_xl")
	defer MyRecover()

	client := &http.Client{}
	for i := 1; i <3 ; i++ {
		url := fmt.Sprintf("https://www.kuaidaili.com/free/inha/%d",i)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println(err)
			return nil
		}
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Cache-Control", "max-age=0")
		req.Header.Set("Upgrade-Insecure-Requests", "1")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36")
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
		req.Header.Set("Sec-Fetch-Site", "same-origin")
		req.Header.Set("Sec-Fetch-Mode", "navigate")
		req.Header.Set("Sec-Fetch-User", "?1")
		req.Header.Set("Sec-Fetch-Dest", "document")
		req.Header.Set("Referer", "https://www.kuaidaili.com/free/inha/2/")
		req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
		req.Header.Set("Cookie", "_ga=GA1.2.1733857456.1593502741; channelid=0; sid=1596347512066172; _gid=GA1.2.1423723694.1596347512; _gat=1")
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			return nil
		}
		bodyText, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return nil

		}

		doc := htmlquery.ReadFromString(string(bodyText))

		list := htmlquery.QueryAll(doc, "//*[@id=\"list\"]/table/tbody/tr")

		for _, n := range list {
			info := proxy.IPInfo{
				IP:     htmlquery.FindOneText(n, "//td[1]"),
				Port:   htmlquery.FindOneText(n, "//td[2]"),
				IPType: switchIPType(htmlquery.FindOneText(n, "//td[4]")),
			}
			infoList = append(infoList, info)
		}
		time.Sleep(time.Second*1)
	}

	return infoList

}
