package spiderProject

import (
	"io/ioutil"
	"log"
	"net/http"
	"proxy/internal/proxy"
	"proxy/pkg/htmlquery"
	"regexp"
)

func switchIPType(t string) int {
	switch t {
	case "HTTP代理":
		return proxy.HTTP
	case "HTTP":
		return proxy.HTTP
	case "HTTPS":
		return proxy.HTTPS
	case "HTTPS代理":
		return proxy.HTTPS
	default:
		return proxy.HTTP
	}
}

func Spider_nimadaili() (infoList []proxy.IPInfo) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://www.nimadaili.com/gaoni/", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Referer", "http://www.nimadaili.com/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", "csrftoken=wqBp9IoBJ6sx8ve8xLXxCkw9OAq7mwTgiO30yi8YLeVWaj0P5xSLeBxe162xMRLs")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc := htmlquery.ReadFromString(string(bodyText))
	tr := htmlquery.QueryAll(doc, "/html/body/div[1]/div[1]/div/table/tbody/tr")
	for _, n := range tr {

		ipMatch, _ := regexp.Compile("((25[0-5]|2[0-4]\\d|1\\d{2}|[1-9]?\\d)\\.){3}(25[0-5]|2[0-4]\\d|1\\d{2}|[1-9]?\\d)")
		portMatch, _ := regexp.Compile(":.*")

		var info = proxy.IPInfo{
			IP:     ipMatch.FindString(htmlquery.FindOneText(n, "//td[1]")),
			Port:   portMatch.FindString(htmlquery.FindOneText(n, "//td[1]"))[1:],
			IPType: switchIPType(htmlquery.FindOneText(n, "//td[2]")),
			Rating: 50,
		}
		infoList = append(infoList, info)
	}
	return infoList

}
