package test2
//
//import (
//	"bufio"
//	"bytes"
//	"fmt"
//	"io"
//	"log"
//	"net"
//	"net/http"
//	"net/url"
//	"os"
//	"runtime/debug"
//	"strings"
//	"sync"
//	"time"
//
//	"github.com/robfig/cron"
//)
//
//var proxyUrls map[string]string = make(map[string]string)
//var choiseURL string
//var mu sync.Mutex
//var connHold map[string]net.Conn = make(map[string]net.Conn) //map[代理服务器url]tcp连接
//
//func init() {
//	log.SetFlags(log.LstdFlags | log.Lshortfile)
//	refreshProxyAddr()
//
//	cronTask := cron.New()
//	cronTask.AddFunc("@every 1h", func() {
//		mu.Lock()
//		defer mu.Unlock()
//		refreshProxyAddr()
//	})
//	cronTask.Start()
//}
//
//func main() {
//	l, err := net.Listen("tcp", ":7856")
//	if err != nil {
//		log.Panic(err)
//	}
//
//	for {
//		client, err := l.Accept()
//		if err != nil {
//			log.Panic(err)
//		}
//		go handle(client)
//	}
//}
//
//func handle(client net.Conn) {
//	defer func() {
//		if err := recover(); err != nil {
//			log.Println(err)
//			debug.PrintStack()
//		}
//	}()
//	if client == nil {
//		return
//	}
//	log.Println("client tcp tunnel connection:", client.LocalAddr().String(), "->", client.RemoteAddr().String())
//	// client.SetDeadline(time.Now().Add(time.Duration(10) * time.Second))
//	defer client.Close()
//
//	var b [1024]byte
//	n, err := client.Read(b[:]) //读取应用层的所有数据
//	if err != nil || bytes.IndexByte(b[:], '\n') == -1 {
//		log.Println(err) //传输层的连接是没有应用层的内容 比如：net.Dial()
//		return
//	}
//	var method, host, address string
//	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
//	log.Println(method, host)
//	hostPortURL, err := url.Parse(host)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	if hostPortURL.Opaque == "443" { //https访问
//		address = hostPortURL.Scheme + ":443"
//	} else { //http访问
//		if strings.Index(hostPortURL.Host, ":") == -1 { //host不带端口， 默认80
//			address = hostPortURL.Host + ":80"
//		} else {
//			address = hostPortURL.Host
//		}
//	}
//
//	server, err := Dial("tcp", address)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	//在应用层完成数据转发后，关闭传输层的通道
//	defer server.Close()
//	log.Println("server tcp tunnel connection:", server.LocalAddr().String(), "->", server.RemoteAddr().String())
//	// server.SetDeadline(time.Now().Add(time.Duration(10) * time.Second))
//
//	if method == "CONNECT" {
//		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
//	} else {
//		log.Println("server write", method) //其它协议
//		server.Write(b[:n])
//	}
//
//	//进行转发
//	go func() {
//		io.Copy(server, client)
//	}()
//	io.Copy(client, server) //阻塞转发
//}
//
////refreshProxyAddr 刷新代理ip
//func refreshProxyAddr() {
//	var proxyUrlsTmp map[string]string = make(map[string]string)
//	//获取代理ip地址逻辑
//	proxyUrls = proxyUrlsTmp //可以手动设置测试代理ip
//}
//
////DialSimple 直接通过发送数据报与二级代理服务器建立连接
//func DialSimple(network, addr string) (net.Conn, error) {
//	var proxyAddr string
//	for proxyAddr = range proxyUrls { //随机获取一个代理地址
//		break
//	}
//	c, err := func() (net.Conn, error) {
//		u, _ := url.Parse(proxyAddr)
//		log.Println("代理host", u.Host)
//		// Dial and create client connection.
//		c, err := net.DialTimeout("tcp", u.Host, time.Second*5)
//		if err != nil {
//			log.Println(err)
//			return nil, err
//		}
//		_, err = c.Write([]byte("CONNECT w.xxxx.com:443 HTTP/1.1\r\n Host: w.xxxx.com:443\r\n User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.3\r\n\r\n"))// w.xxxx.com:443 替换成实际的地址
//		if err != nil {
//			panic(err)
//		}
//		c.Write([]byte(`GET www.baidu.com HTTP/1.1\r\n\r\n`))
//		io.Copy(os.Stdout, c)
//		return c, err
//	}()
//	return c, err
//}
//
////Dial 建立一个传输通道
//func Dial(network, addr string) (net.Conn, error) {
//	var proxyAddr string
//	for proxyAddr = range proxyUrls { //随机获取一个代理地址
//		break
//	}
//	//建立到代理服务器的传输层通道
//	c, err := func() (net.Conn, error) {
//		u, _ := url.Parse(proxyAddr)
//		log.Println("代理地址", u.Host)
//		// Dial and create client connection.
//		c, err := net.DialTimeout("tcp", u.Host, time.Second*5)
//		if err != nil {
//			return nil, err
//		}
//
//		reqURL, err := url.Parse("http://" + addr)
//		if err != nil {
//			return nil, err
//		}
//		req, err := http.NewRequest(http.MethodConnect, reqURL.String(), nil)
//		if err != nil {
//			return nil, err
//		}
//		req.Close = false
//		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.3")
//
//		err = req.Write(c)
//		if err != nil {
//			return nil, err
//		}
//
//		resp, err := http.ReadResponse(bufio.NewReader(c), req)
//		if err != nil {
//			return nil, err
//		}
//		defer resp.Body.Close()
//
//		log.Println(resp.StatusCode, resp.Status, resp.Proto, resp.Header)
//		if resp.StatusCode != 200 {
//			err = fmt.Errorf("Connect server using proxy error, StatusCode [%d]", resp.StatusCode)
//			return nil, err
//		}
//		return c, err
//	}()
//	if c == nil || err != nil { //代理异常
//		log.Println("代理异常：", c, err)
//		log.Println("本地直接转发：", c, err)
//		return net.Dial(network, addr)
//	}
//	log.Println("代理正常,tunnel信息", c.LocalAddr().String(), "->", c.RemoteAddr().String())
//	return c, err
//}