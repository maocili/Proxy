package service

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"strings"
	"time"
)

func handleClientRequest(conn net.Conn) {
	if conn == nil {
		return
	}
	defer conn.Close()

	var b [1024]byte
	n, err := conn.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b[:]))

	var method, host, address string
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
	hostPortURL, err := url.Parse(host)
	if err != nil {
		log.Println(err)
		return
	}
	if hostPortURL.Opaque == "443" { //https访问
		address = hostPortURL.Scheme + ":443"
	} else { //http访问
		if strings.Index(hostPortURL.Host, ":") == -1 { //host不带端口， 默认80
			address = hostPortURL.Host + ":80"
		} else {
			address = hostPortURL.Host
		}
	}

	//获得了请求的host和port，就开始拨号吧
	server, err := Dial()
	fmt.Println(address)
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Close()

	log.Println("server tcp tunnel connection:", server.LocalAddr().String(), "->", server.RemoteAddr().String())
	// server.SetDeadline(time.Now().Add(time.Duration(10) * time.Second))

	if method == "CONNECT" {
		fmt.Fprint(conn, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		log.Println("server write", method) //其它协议
		server.Write(b[:n])
	}

	go func() {
		io.Copy(server, conn)
	}()
	io.Copy(conn, server) //阻塞转发

}

//直接转发代理
func Dial() (net.Conn, error) {
	var proxyAddr string
	proxyAddr = pool.RandIP()

	c, err := func() (net.Conn, error) {
		u, _ := url.Parse(proxyAddr)
		log.Println("代理host", u.Host)
		// Dial and create client connection.
		c, err := net.DialTimeout("tcp", u.Host, time.Second*5)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return c, err
	}()
	return c, err
}
