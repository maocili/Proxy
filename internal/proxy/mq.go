package proxy

import "fmt"

type MQ struct {
	ch       chan IPInfo //用于接收新ipd的chan
	Consumer func(ch <-chan IPInfo, pool *ProxyPool)
}

func NewMQ(buffer int) *MQ {

	return &MQ{
		ch:       make(chan IPInfo, buffer),
		Consumer: InsertIP,
	}

}

// 消费者
func InsertIP(i <-chan IPInfo, pool *ProxyPool) {
	for ipinfo := range i {
		info := ipinfo
		ip := fmt.Sprint(info)
		pool.Append(ProxyIP(ip), info)
	}

}
