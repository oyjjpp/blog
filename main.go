package main

import (
	"context"
	"fmt"
	"github.com/oyjjpp/blog/service/queue"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/CodeLineage/brotli"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/oyjjpp/blog/global"
	"github.com/oyjjpp/blog/initialize"
	"github.com/oyjjpp/blog/middleware"
	"github.com/oyjjpp/blog/route"
)

func initDb() {
	// 初始化数据库
	initialize.Mysql()
	initialize.DBTables()

	// 程序结束前关闭数据库链接
	defer func() {
		if sqlDB, err := global.MysqlDB.DB(); err == nil {
			log.Println("程序结束前关闭数据库链接")
			sqlDB.Close()
		}
	}()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// 初始化kafka
	log.Println("init kafka")
	queue.ProducerInit(ctx)
	// 注册HTTP服务
	ginCreate()

	// 信号
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	<-exit
	cancel()
	queue.CloseKafka()
}

// ginCreate
func ginCreate() {
	handler := gin.New()

	// 设置运行模式
	gin.SetMode(gin.DebugMode)
	gin.ForceConsoleColor()

	// 注册中间件
	handler.Use(middleware.Logger())
	// handler.Use(gin.Recovery())

	// 注册gzip中间件
	// handler.Use(gzip.Gzip(gzip.DefaultCompression))
	// handler.Use(gzip2.DefaultHandler().Gin)
	handler.Use(brotli.NewHandler(brotli.Config{
		CompressionLevel: brotli.DefaultCompression,
		MinContentLength: brotli.DefalutContentLen,
		RequestFilter: []brotli.RequestFilter{
			brotli.NewCommonRequestFilter(),
			brotli.NewRequestApiFilter([]string{
				"/blog/cache/http",
				"/blog/detail/index",
			}),
		},
		ResponseHeaderFilter: []brotli.ResponseHeaderFilter{
			brotli.DefaultContentTypeFilter(),
		},
	}).Gin)

	// 注册路由
	route.LoadRoute(handler)

	// 注册性能分析
	// pprof.Register(handler)

	// 监听端口
	handler.Run(":80") // 监听并在 0.0.0.0:8080 上启动服务
}

// endlessCreate
func endlessCreate() {
	// 设置服务属性
	endless.DefaultReadTimeOut = 5 * time.Second
	endless.DefaultWriteTimeOut = 5 * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", 8091)

	// 创建服务
	server := endless.NewServer(endPoint, initRouter())

	server.BeforeBegin = func(add string) {
		log.Printf("actual pid is %d", syscall.Getpid())
	}

	// 开启服务监听
	if err := server.ListenAndServe(); err != nil {
		log.Printf("server err :%v", err)
	}
}

// serverCreate
// 通过server Shutdown
func serverCreate() {
	router := initRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8091),
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// 通过协程监听服务
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Listen:%s \n", err)
		}
	}()

	// TODO 如何接收信号
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
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
