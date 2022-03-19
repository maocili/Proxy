package proxy

import (
	"sync"
	"time"
)

type (
	ProxyIP string
)

const (
	HTTP  = 1
	HTTPS = 2
	SOCKS = 3
)

type IPInfo struct {
	IP     string
	Port   string
	IPType int           //IP类型
	Rating int8          //ip评分
	Alive  time.Duration //存活时间
}

type ProxyPool struct {
	m      sync.RWMutex
	ticker <-chan time.Time //代理池检查周期
	d      time.Duration    //代理池检查周期，仅用于string输出
	ips    map[ProxyIP]IPInfo
}
