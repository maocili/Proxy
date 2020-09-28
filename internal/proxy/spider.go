package proxy

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type SpiderWorker struct {
	m      sync.RWMutex
	Ticker <-chan time.Time
	d      time.Duration //worker 执行时间
	Work   func() []IPInfo
}

//创建一个定时爬虫工作
func NewSpider(d time.Duration, work func() []IPInfo) *SpiderWorker {
	return &SpiderWorker{
		Ticker: time.Tick(d),
		d:      d,
		Work:   work,
	}

}

//启动定时爬虫工作
func (s *SpiderWorker) Start(pool *ProxyPool) {
	log.Println("SpiderWorker Start")
	go func() {
		for {
			select {
			case <-s.Ticker:
				q := NewInsertQueue(64,insertIP)
				go q.Consumer(q.ch, pool)
				list := s.Work()
				//生成者
				for _, info := range list {
					ipInfo := info
					q.ch <- ipInfo
				}}

		}
	}()
}


// 队列Consumer func :插入ip
func insertIP(i <-chan IPInfo, pool *ProxyPool) {
	for ipinfo := range i {
		info := ipinfo
		ip := fmt.Sprint(info)
		pool.Append(ProxyIP(ip), info)
	}

}