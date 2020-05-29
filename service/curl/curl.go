package curl

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"runtime"
	"time"
)

func createDial(network, addr string) (net.Conn, error) {
	dial := net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	conn, err := dial.Dial(network, addr)
	if err != nil {
		return conn, err
	}
	fmt.Println("connect done use", conn.LocalAddr().String())
	return conn, err
}

func createClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: createDial,
		},
	}
	return client
}

func createTransport(localAddr net.Addr) *http.Transport {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}
	if localAddr != nil {
		dialer.LocalAddr = localAddr
	}
	return &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
	}
}

func doGet(url string, id int) {
	res, err := createClient().Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	buf, err := ioutil.ReadAll(res.Body)
	fmt.Printf("%d:%s -- %v\n", id, string(buf), err)
	if err := res.Body.Close(); err != nil {
		fmt.Println(err.Error())
	}
}
