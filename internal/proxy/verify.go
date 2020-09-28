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
				q := NewInsertQueue(64, ratingIP)
				go q.Consumer(q.ch, pool)
				pool.VerifyLoop(q)
			}
		}
	}()
}

func ratingIP(i <-chan IPInfo, pool *ProxyPool) {
	for ipinfo := range i {
		info := ipinfo
		ip := fmt.Sprint(info)
		go pool.Alter(ProxyIP(ip), info)
	}
}

func (pool *ProxyPool) VerifyLoop(q *Queue) {

	var wg sync.WaitGroup
	it := reflect.ValueOf(pool.ips).MapRange()

	//ips loop
	for it.Next() {
		pool.m.Lock()
		wg.Add(1)
		go verifyIP(&wg, q, pool.ips[ProxyIP(it.Key().String())])
		pool.m.Unlock()
	}
	wg.Wait()
}

func verifyIP(wg *sync.WaitGroup, q *Queue, info IPInfo) {
	defer wg.Done()
	//fmt.Println("verifyIP:", info)

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(info.String())
	}

	transport := &http.Transport{Proxy: proxy}

	client := &http.Client{Transport: transport}
	resp, err := client.Get("http://116.62.125.79/get")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		info.Rating -= 30
		q.ch <- info
	}

}
