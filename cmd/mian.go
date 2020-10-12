package main

import (
	"proxy/internal/api"
	"proxy/internal/proxy"
	"proxy/spiderProject"
	"time"
)

func work() []proxy.IPInfo {
	var infoList []proxy.IPInfo

	info := proxy.IPInfo{
		IP:   "132132",
		Port: "123",
	}
	infoList = append(infoList, info)
	info = proxy.IPInfo{
		IP:   "789789",
		Port: "1778978923",
	}
	infoList = append(infoList, info)

	return infoList
}

func main() {
	//test2.M()

	p := proxy.NewPool(time.Second * 30)
	p.StartVerify()
	s := proxy.NewSpider(time.Second*60, spiderProject.Spider_kuaidaili)
	s.Start(p)
	nimadaili := proxy.NewSpider(time.Second*30, spiderProject.Spider_nimadaili)
	nimadaili.Start(p)

	api.StartWebServe(p,":8080")
	defer func() { select {} }()

}
