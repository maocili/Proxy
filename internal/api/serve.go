package api

import (
	"github.com/gin-gonic/gin"
	"proxy/internal/proxy"
)

var pool *proxy.ProxyPool

func StartWebServe(p *proxy.ProxyPool, addr string) {
	pool = p
	router := gin.Default()
	router.GET("/list", GetList)
	router.GET("/rand", RandIP)
	router.Run(addr)
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
