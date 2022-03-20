package http

import (
	"github.com/gin-gonic/gin"
	"proxy/internal/http/font"
	"proxy/internal/http/service"
)

func StartWebServe(addr string) {

	router := gin.Default()

	router.GET("/list", http.GetList)
	router.GET("/rand", http.RandIP)
	router.POST("/append", http.AppendIP)
	router.POST("/delete", http.DeleteIP)

	router.Static("/assets", "internal/http/font/templates/assets")

	router.LoadHTMLGlob("internal/http/font/templates/page/*")
	web := router.Group("/web")
	web.GET("/index", font.Index)

	router.Run(addr)

}
