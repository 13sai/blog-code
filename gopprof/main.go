package main

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	var debugHttp *http.Server

	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(gin.Logger())

	pprof.Register(g)

	// 使用9000端口开启http服务
	debugHttp = &http.Server{
		Addr:    ":9000",
		Handler: g,
	}
	debugHttp.ListenAndServe()
}
