package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/oyjjpp/blog/route"
)

func main() {
	go func() {
		handler := gin.New()
		route.LoadRoute(handler)
		pprof.Register(handler)

		handler.Run() // 监听并在 0.0.0.0:8080 上启动服务
	}()

	go func() {
		fmt.Println("申请内存1")
		tick := time.NewTicker(time.Second * 5)
		var buf []byte
		fmt.Println("申请内2")
		for range tick.C {
			fmt.Println("申请内存3")
			if len(buf) > 1024*1024*100 {
				tick.Stop()
			}
			fmt.Println("申请内存4")
			buf = append(buf, make([]byte, 1024*1024)...)
		}
	}()
	time.Sleep(time.Minute * 10)
}
