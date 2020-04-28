package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/oyjjpp/blog/routes"
)

func main() {
	route := gin.New()
	routes.LoadRoute(route)
	pprof.Register(route)
	route.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
