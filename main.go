package main

import (
	"fmt"
	"log"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/oyjjpp/blog/route"
)

func main() {
	endlessCreate()
}

// ginCreate
func ginCreate() {
	endless.DefaultReadTimeOut = 5
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

// endlessCreate
func endlessCreate() {
	endless.DefaultReadTimeOut = 5 * time.Second
	endless.DefaultWriteTimeOut = 5 * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", 8080)

	server := endless.NewServer(endPoint, initRouter())

	server.BeforeBegin = func(add string) {
		log.Printf("actual pid is %d", syscall.Getpid())
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("server err :%v", err)
	}
}

// initRouter
func initRouter() *gin.Engine {
	handler := gin.New()
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	// 注册路由
	route.LoadRoute(handler)
	return handler
}
