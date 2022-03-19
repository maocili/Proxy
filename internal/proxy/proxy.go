package proxy

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

//创建Pool
func NewPool(d time.Duration) *ProxyPool {
	return &ProxyPool{
		ticker: time.Tick(d),
		d:      d,
		ips:    make(map[ProxyIP]IPInfo),
	}
}

//添加一个新ip,如果存在则会跳转至Alter()
func (p *ProxyPool) Append(ip ProxyIP, info IPInfo) {

	p.m.Lock()

	if _, isexist := p.ips[ip]; isexist {
		// TODO: 存在时？
		p.m.Unlock()
		//p.Alter(ip, info)
		return
	}
	p.ips[ip] = info
	p.m.Unlock()

	// TODO: 添加新ip时先进行检测

}

func (p *ProxyPool) Alter(ip ProxyIP, info IPInfo) {

	p.m.Lock()
	defer p.m.Unlock()

	if _, isexist := p.ips[ip]; isexist {
		p.ips[ip] = info
	}
}

func (p *ProxyPool) Delete(ip ProxyIP) {

	p.m.Lock()
	defer p.m.Unlock()

	delete(p.ips, ip)
}

//返回一个ip
func (p *ProxyPool) RandIP() string {

	p.m.Lock()
	defer p.m.Unlock()

	for ip, info := range p.ips {
		if info.Rating >= 60 {
			return string(ip)
		}
	}
	return ""

}

//获取所有的代理ip
func (p *ProxyPool) GetList() (data []IPInfo) {
	p.m.RLock()
	defer p.m.RUnlock()
	ipIter := reflect.ValueOf(p.ips).MapRange()
	for ipIter.Next() {
		info := p.ips[ProxyIP(ipIter.Key().String())]
		//data = append(data, fmt.Sprintf("ip: %s |rating: %d", info, info.Rating))
		info.IP = info.String()
		data = append(data, info)

	}
	return data

}

//获取n个ip
func (p *ProxyPool) GetNIP(n int) (data []string) {
	p.m.Lock()
	defer p.m.Unlock()
	ipIter := reflect.ValueOf(p.ips).MapRange()
	for i := 0; i <= n; i++ {
		if ipIter.Next() {
			data = append(data, ipIter.Key().String())
		}
	}

	return data
}

func (p ProxyPool) String() string {
	b, _ := json.Marshal(p.ips)
	return fmt.Sprintf("检查周期：%d 秒 \n"+
		"i p 池 : %s \n", p.d/time.Second, string(b))
}

func (i IPInfo) String() string {
	switch i.IPType {
	case HTTP:
		return fmt.Sprintf("http://%s:%s", i.IP, i.Port)
	case HTTPS:
		return fmt.Sprintf("https://%s:%s", i.IP, i.Port)
	}
	return fmt.Sprintf("%d://%s:%s", i.IPType, i.IP, i.Port)
}

func (info IPInfo) Host() string {
	return info.IP + ":" + info.Port
}
