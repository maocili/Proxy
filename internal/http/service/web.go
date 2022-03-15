package http

import (
	"github.com/gin-gonic/gin"
	"proxy/internal/proxy/service"
)

var proxyService service.ProxyService

func GetList(context *gin.Context) {

	context.JSON(200, gin.H{
		"ip": proxyService.GetList(),
	})
	context.Abort()

}

func RandIP(c *gin.Context) {

	c.JSON(200, gin.H{
		"ip": proxyService.RandIP(),
	})
	c.Abort()

}
