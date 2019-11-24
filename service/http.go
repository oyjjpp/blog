//发送HTTP请求
package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

//发送一个http请求
func HttpGet(url string) string {
	//生成client 参数为默认
	client := &http.Client{}

	//第三个参数是io.reader interface
	//strings.NewReader byte.NewReader bytes.NewBuffer  实现了read 方法
	// 如果存在参数，只需序列化成字符串，并转化为io.reader类型，传入第三个参数
	// 举个例子：?name=pj&age=18 则传入参数可以为: strings.NewReader("name=pj&age=18")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle err
	}
	// 设置请求头
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//处理返回结果
	response, err := client.Do(req)
	if err != nil {

	}
	//结束后关闭回复的主体
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// handle err
	}
	return string(body)
}

// 复杂的请求参数
func HttpPost(url string) string {
	type Server struct {
		ServerName string
		ServerIp   string
		Age        int
	}

	type ServerSlice struct {
		Server    []Server
		ServersID string
	}
	//post 第三个参数是io.reader interface
	//strings.NewReader byte.NewReader bytes.NewBuffer  实现了read 方法
	s := ServerSlice{
		ServersID: "tearm",
		Server:    []Server{{"beijing", "127.0.0.1", 20}, {"shanghai", "127.0.0.1", 22}},
	}
	b, _ := json.Marshal(s) // 这里的s既可以是结构体，也可以是map[][]类型

	client := &http.Client{}
	// 如果存在参数，只需序列化成字符串，并转化为io.reader类型，传入第三个参数
	req, err := http.NewRequest("POST", url, strings.NewReader("heel="+string(b)))
	if err != nil {
		panic(err)
	}
	// 设置请求数据类型："application/json"，"application/x-www-form-urlencoded"
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
