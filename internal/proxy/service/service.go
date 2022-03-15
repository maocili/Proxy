package service

import (
	"log"
	"net"
	"proxy/internal/proxy"
)

var pool *proxy.ProxyPool

type ProxyService struct {
}

func Proxy(p *proxy.ProxyPool) {
	pool = p
	log.Println("http proxy service running port :10088")
	l, err := net.Listen("tcp", ":10088")
	if err != nil {
		log.Panic(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleClientRequest(conn)
	}
}

func (p ProxyService) Append(ip proxy.ProxyIP, info proxy.IPInfo) {
	pool.Append(ip, info)
}

func (p ProxyService) Alter(ip proxy.ProxyIP, info proxy.IPInfo) {
	pool.Alter(ip, info)
}

func (p ProxyService) Delete(ip proxy.ProxyIP) {
	pool.Delete(ip)
}

func (p ProxyService) GetList() []proxy.IPInfo {
	return pool.GetList()
}

func (p ProxyService) GetNIP(n int) []string {
	return pool.GetNIP(n)
}

func (p ProxyService) RandIP() string {
	return pool.RandIP()
}
