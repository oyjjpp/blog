//发送HTTP请求
package service

import (
	"blog/service/httpclient"
	"blog/util"
	"net/url"
	"strings"
	"time"
)

/*
//发送一个http请求
func HttpGetV1(url string) string {
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
func HttpPostV1(url string) string {
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
*/
//HTTPClient 用于发出http 请求
type HTTPClient struct {
	client *httpclient.HttpClient
}

//HTTPArgs 请求调用传入的参数
type HTTPArgs struct {
	//路径
	Path string
	//扩展参数
	Params map[string]string

//Get http get方法
func (http *HTTPClient) Get(args HTTPArgs) (string, error) {

	//配置timeout
	timeout, err := util.Int(args.Params["timeout"])
	if err != nil {
		timeout = 1
	}
	http.parseTimeout(timeout)
	delete(args.Params, "timeout")

	//请求地址
	url := args.Path
	//设置代理
	if proxyURL := args.Params["proxyURL"]; proxyURL != "" {
		args.Header["proxyUrl"] = proxyURL
		delete(args.Params, "proxyURL")
	}

	//响应结果
	var res *httpclient.Response

	//重试次数
	tryCount, err := util.Int(args.Params["TryCount"])
	if err != nil {
		tryCount = 1
	}
	delete(args.Params, "TryCount")

	//发送请求
	for i := 0; i < tryCount; i++ {
		res, err = http.client.Get(url, args.Params, args.Header)
		if err == nil {
			break
		}
	}

	var jsonStr string
	if err == nil {
		jsonStr, err = res.ToString()
	}

	return jsonStr, err
}

//解析处理http timeout参数
func (http *HTTPClient) parseTimeout(timeout int) {
	if timeout != 0 {
		timeOut := time.Duration(timeout * 1000)
		http.client.SetTimeOut(time.Millisecond * timeOut)
	}
}

//把param参数追加到url后面
func (http *HTTPClient) urlAppendParam(url string, param map[string]string) string {
	str := paramsToString(param)
	if str == "" {
		return url
	}

	if strings.IndexByte(url, '?') == -1 {
		url += "?" + str
	} else {
		url += "&" + str
	}

	return url
}

//对param编码为url编码格式("bar=baz&foo=quux"
func paramsToString(params map[string]string) string {
	values := url.Values{}
	for k, v := range params {
		values.Set(k, v)
	}

	return values.Encode()
}
