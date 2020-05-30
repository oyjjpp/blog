package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/oyjjpp/blog/route"
)

func main() {
	handler := gin.New()
	// 注册路由
	route.LoadRoute(handler)
	// 注册性能分析
	pprof.Register(handler)
	// 注册中间件
	handler.Use(gin.Logger())
	// 监听端口
	handler.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
