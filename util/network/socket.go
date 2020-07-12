package network

import (
	"net"
	"net/http"
	"runtime"
	"time"
)

// 全局公用的transport方便于做长连接和连接池
var transport *http.Transport

func init() {
	transport = createTransport()
}

func newClient() *http.Client {
	return &http.Client{
		Transport: transport,
	}
}

func createTransport() *http.Transport {
	dialer := &net.Dialer{
		Timeout:   1 * time.Second,
		KeepAlive: 3 * time.Second,
	}
	return &http.Transport{
		DialContext:         dialer.DialContext,
		MaxIdleConnsPerHost: runtime.GOMAXPROCS(0) + 1,
	}
}
