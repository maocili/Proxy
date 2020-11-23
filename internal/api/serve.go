package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"proxy/internal/proxy"
)

var pool *proxy.ProxyPool

func StartWebServe(p *proxy.ProxyPool, addr string) {
	pool = p
	router := gin.Default()
	router.GET("/list", GetList)
	router.GET("/rand", RandIP)
	router.POST("/add", AddIP)
	router.GET("/sub",Sub)
	router.Run(addr)

}

func Sub(context *gin.Context) {
	pool.RandIP()
}

func GetList(context *gin.Context) {

	context.JSON(200, gin.H{
		"ip": pool.GetList(),
	})
	context.Abort()

}

func RandIP(c *gin.Context) {

	c.JSON(200, gin.H{
		"ip": pool.RandIP(),
	})
	c.Abort()

}

func AddIP(context *gin.Context) {

	var ipinfo proxy.IPInfo

	if err := context.ShouldBindJSON(&ipinfo); err != nil {
		context.Abort()
		fmt.Println(err.Error())
		return
	}

	ipinfo.Rating = 50
	pool.Append(proxy.ProxyIP(ipinfo.String()), ipinfo)

	context.JSON(http.StatusOK, gin.H{
		"msg": "succuess",
	})
	context.Abort()
}

