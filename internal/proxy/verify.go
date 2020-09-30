package proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"sync"
)

func (pool *ProxyPool) StartVerify() {
	log.Println("Start Verify IP Loop")
	go func() {
		for {
			select {
			case <-pool.ticker:
				log.Println("Running Verify")
				q := NewInsertQueue(64, consumerRatingIP)
				go q.Consumer(q.ch, pool)
				pool.VerifyLoop(q)
			}
		}
	}()
}

func consumerRatingIP(i <-chan IPInfo, pool *ProxyPool) {
	for ipinfo := range i {
		info := ipinfo
		ip := ProxyIP(fmt.Sprint(info))
		if info.Rating <= 0 {
			go pool.Delete(ip) //TODO:防止源地址提供重复的ip
		} else {
			go pool.Alter(ip, info)
		}
	}
}

func (pool *ProxyPool) VerifyLoop(q *Queue) {

	var wg sync.WaitGroup
	it := reflect.ValueOf(pool.ips).MapRange()

	//ips loop
	for it.Next() {
		pool.m.Lock()
		wg.Add(1)
		go func(wg *sync.WaitGroup, queue *Queue, info IPInfo) {

			defer wg.Done()

			if verifyIP(info) {
				info.Rating += 10
			} else {
				info.Rating -= 30
			}
			q.ch <- info

		}(&wg, q, pool.ips[ProxyIP(it.Key().String())])
		pool.m.Unlock()
	}
	wg.Wait()

}

func verifyIP(info IPInfo) bool {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(info.String())
	}

	transport := &http.Transport{Proxy: proxy}

	client := &http.Client{Transport: transport}
	resp, err := client.Get("http://116.62.125.79/get")
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false
	} else {
		return true
	}
}
