package http

import (
	"github.com/gin-gonic/gin"
	"proxy/internal/http/service"
)

func StartWebServe(addr string) {

	router := gin.Default()
	router.GET("/list", http.GetList)
	router.GET("/rand", http.RandIP)
	router.POST("/append", http.AppendIP)
	router.POST("/delete", http.DeleteIP)
	router.Run(addr)

}
