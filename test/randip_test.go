package test

import (
	"os"
	proxy "proxy/internal/proxy"
	"strings"
	"testing"
	"time"
)

func tracefile(str_content string)  {
	fd,_:=os.OpenFile("a.txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	fd_content:=strings.Join([]string{str_content,"\n"},"")
	buf:=[]byte(fd_content)
	fd.Write(buf)
	fd.Close()
}

func BenchmarkProxy(b *testing.B) {
	p := proxy.NewPool(time.Second*3)

	info := proxy.IPInfo{
		IPType: proxy.HTTP,
		Rating: 50,
	}
	p.Append(proxy.ProxyIP("192.168.0.1"),info)
	p.Append(proxy.ProxyIP("192.168.0.2"),info)
	p.Append(proxy.ProxyIP("192.168.0.3"),info)
	p.Append(proxy.ProxyIP("192.168.0.4"),info)
	p.Append(proxy.ProxyIP("192.168.0.5"),info)
	p.Append(proxy.ProxyIP("192.168.0.6"),info)
	p.Append(proxy.ProxyIP("192.168.0.7"),info)
	p.Append(proxy.ProxyIP("192.168.0.8"),info)
	p.Append(proxy.ProxyIP("192.168.0.9"),info)
	p.Append(proxy.ProxyIP("192.168.0.0"),info)


	b.ResetTimer()
	for i := 0; i < 10000; i++ {
		//fmt.Println(p.RandIP())
		tracefile(p.RandIP())
	}


}
