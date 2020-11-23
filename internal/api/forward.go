package api

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"proxy/filter"
	"proxy/internal/proxy"
	"strings"
	"time"
)

type ProxyServe struct {
	Address string
	SubType interface{}
}

func NewProxyListen(add string, sub interface{}) *ProxyServe {

	return &ProxyServe{
		Address: add, //监听端口
		SubType: sub, //订阅类型
	}
}

func (p *ProxyServe) StartProxy() {
	log.Println("Start Proxy Serve")
	//:TODO 过滤ip 类型	p.SubType
	go forward(p)
}

func forward(p *ProxyServe) {
	listen, err := net.Listen("tcp", p.Address)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Listening and proxying on ", p.Address)
	defer listen.Close()

	for {
		client, err := listen.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handle(client)
	}
}

func handle(client net.Conn) {

	//捕获异常
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	if client == nil {
		return
	}
	log.Println("client tcp tunnel connection:", client.LocalAddr().String(), "->", client.RemoteAddr().String())
	defer client.Close()

	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil || bytes.IndexByte(b[:], '\n') == -1 {
		log.Print(err)
		return
	}

	fmt.Print(string(b[:n]))
	var method, host, address string
	var subType int
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
	log.Println(method, host)

	hostPortURL, err := url.Parse(host)
	if err != nil {
		log.Println(err)
		return
	}

	if hostPortURL.Opaque == "443" {
		address = hostPortURL.Scheme + ":443"
		subType = proxy.HTTPS
	} else {
		if strings.Index(hostPortURL.Host, ":") == -1 { // host不带端口号，默认80
			address = hostPortURL.Host + ":80"
		} else {
			address = hostPortURL.Host
		}
		subType = proxy.HTTP
	}

	//设置filterFunc
	var filterFunc func(info proxy.IPInfo) bool

	switch subType {
	case proxy.HTTP:
		filterFunc = filter.HttpFilter
	case proxy.HTTPS:
		filterFunc = filter.HttpsFilter
	default:
		filterFunc = filter.HttpFilter

	}

	server, err := dial("tcp", address, filterFunc)
	if err != nil {
		log.Println(err)
		return
	}

	defer server.Close()
	log.Println("server tcp tunnel connection:", server.LocalAddr().String(), "->", server.RemoteAddr().String())

	if method == "CONNECT" {
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		log.Println("server write", method) //其它协议
		server.Write(b[:n])
	}

	//进行转发
	go func() {
		io.Copy(server, client)
	}()
	io.Copy(client, server) //阻塞转发
}

func dial(network string, address string, filterFunc func(info proxy.IPInfo) bool) (net.Conn, error) {

	// It extracts IP address that according to the filter
	proxyAddr := pool.FilterIP(filterFunc)

	conn, err := func() (net.Conn, error) {
		log.Println("代理地址", proxyAddr.Host())

		conn, err := net.DialTimeout("tcp", proxyAddr.Host(), time.Second*15)
		if err != nil {
			return nil, err
		}

		reqURL, err := url.Parse("http://" + address)
		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest(http.MethodConnect, reqURL.String(), nil)
		if err != nil {
			return nil, err
		}

		err = req.Write(conn)
		if err != nil {
			return nil, err
		}

		res, err := http.ReadResponse(bufio.NewReader(conn), req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		log.Println(res.StatusCode, res.Status, res.Proto, res.Header)
		if res.StatusCode != 200 {
			err = fmt.Errorf("代理错误, StatusCode [%d]", res.StatusCode)
			return nil, err
			//:TODO 重试IP -> 更换代理IP
		}
		return conn, err
	}()

	if conn == nil || err != nil {
		//:TODO 重试IP -> 更换代理IP
		log.Println("代理异常：", conn, err)
		log.Println("本地直接转发：", conn, err)
		return net.Dial(network, address)

	}
	log.Println("代理正常,tunnel信息", conn.LocalAddr().String(), "->", conn.RemoteAddr().String())
	return conn, err

}
