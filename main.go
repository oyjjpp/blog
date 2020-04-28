package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/oyjjpp/blog/route"
)

func main() {
	handler := gin.New()
	route.LoadRoute(handler)
	pprof.Register(handler)
	handler.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
