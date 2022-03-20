package http

import (
	"github.com/gin-gonic/gin"
	proxy2 "proxy/internal/proxy"
	"proxy/internal/proxy/service"
)

var proxyService service.ProxyService

func GetList(c *gin.Context) {

	c.JSON(200, gin.H{
		"ip": proxyService.GetList(),
	})
	c.Abort()

}

func RandIP(c *gin.Context) {

	c.JSON(200, gin.H{
		"ip": proxyService.RandIP(),
	})
	c.Abort()

}

func AppendIP(c *gin.Context) {

	var ipinfo IPInfo

	if err := c.ShouldBindJSON(&ipinfo); err != nil {
		c.JSON(304, gin.H{
			"AppendIP err": err.Error(),
		})
		return
	}

	var proxy service.ProxyService
	proxyInfo := proxy2.IPInfo{
		IP:     ipinfo.IP,
		Port:   ipinfo.Port,
		IPType: ipinfo.IPType,
		Rating: 50,
		Alive:  0,
	}
	ip := proxyInfo.String()
	proxy.Append(proxy2.ProxyIP(ip), proxyInfo)
	c.JSON(200, gin.H{
		"msg":  "successfully append ip ",
		"data": ipinfo,
	})
}

func DeleteIP(c *gin.Context) {
	var ipinfo IPInfo
	if err := c.BindHeader(&ipinfo); err != nil {
		c.JSON(304, gin.H{
			"DeleteIP err": err.Error(),
		})
	}

	var proxy service.ProxyService
	proxyInfo := proxy2.IPInfo{
		IP:     ipinfo.IP,
		Port:   ipinfo.Port,
		IPType: ipinfo.IPType,
		Rating: 50,
		Alive:  0,
	}
	proxyIP := proxyInfo.String()
	proxy.Delete(proxy2.ProxyIP(proxyIP))
	c.JSON(200, gin.H{
		"msg":  "successful delete",
		"data": ipinfo,
	})
}
