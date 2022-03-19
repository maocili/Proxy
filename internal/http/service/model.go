package http

import (
	"fmt"
	"proxy/internal/proxy"
)

type IPInfo struct {
	IP     string `json:"ip"`
	Port   string `json:"port"`
	IPType int    `json:"ip_type"`
}

func (p IPInfo) String() string {

	return fmt.Sprintf("%s://%s:%s", p.IPType, p.IP, p.Port)
}

func (ipinfo *IPInfo) ForProxyIpInfo() proxy.IPInfo {

	proxyIPinfo := proxy.IPInfo{
		IP:     ipinfo.IP,
		Port:   ipinfo.Port,
		IPType: ipinfo.IPType,
		Rating: 50,
		Alive:  0,
	}

	return proxyIPinfo

}
