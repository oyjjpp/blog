package httpclient

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"sync"
)

// Constants definations
// CURL options, see https://github.com/bagder/curl/blob/169fedbdce93ecf14befb6e0e1ce6a2d480252a3/packages/OS400/curl.inc.in
const (
	VERSION   = "0.6.2"
	USERAGENT = "go-httpclient v" + VERSION

	PROXY_HTTP    = 0
	PROXY_SOCKS4  = 4
	PROXY_SOCKS5  = 5
	PROXY_SOCKS4A = 6

	// CURL like OPT
	OPT_AUTOREFERER       = 58
	OPT_FOLLOWLOCATION    = 52
	OPT_CONNECTTIMEOUT    = 78
	OPT_CONNECTTIMEOUT_MS = 156
	OPT_MAXREDIRS         = 68
	OPT_PROXYTYPE         = 101
	OPT_TIMEOUT           = 13
	OPT_TIMEOUT_MS        = 155
	OPT_COOKIEJAR         = 10082
	OPT_INTERFACE         = 10062
	OPT_PROXY             = 10004
	OPT_REFERER           = 10016
	OPT_USERAGENT         = 10018

	// Other OPT
	OPT_REDIRECT_POLICY = 100000
	OPT_PROXY_FUNC      = 100001
	OPT_DEBUG           = 100002
	//指定用于tls.Client的TLS配置信息
	OPT_UNSAFE_TLS = 100004
)

// String map of options
var CONST = map[string]int{
	//TRUE时将根据 Location: 重定向时，自动设置header中的Referer:信息
	"OPT_AUTOREFERER": 58,

	//TRUE时将会根据服务器返回HTTP 头中的 "Location: " 重定向
	"OPT_FOLLOWLOCATION": 52,

	//在尝试连接时等待的秒数。设置为0，则无限等待
	"OPT_CONNECTTIMEOUT": 78,

	//尝试连接等待的时间，以毫秒为单位。设置为0，则无限等待
	"OPT_CONNECTTIMEOUT_MS": 156,

	//指定最多的HTTP重定向次数
	"OPT_MAXREDIRS": 68,

	//代理类型
	"OPT_PROXYTYPE": 101,

	//允许cURL函数执行的最长秒数
	"OPT_TIMEOUT": 13,

	//设置cURL允许执行的最长毫秒数
	"OPT_TIMEOUT_MS": 155,

	//连接结束后，比如，调用 curl_close 后，保存 cookie 信息的文件
	"OPT_COOKIEJAR": 10082,

	//发送的网络接口（interface），可以是一个接口名、IP 地址或者是一个主机名。
	"OPT_INTERFACE": 10062,

	//HTTP代理通道
	"OPT_PROXY": 10004,

	//在HTTP请求头中"Referer: "的内容
	"OPT_REFERER": 10016,

	//在HTTP请求中包含一个"User-Agent: "头的字符串
	"OPT_USERAGENT": 10018,

	"OPT_REDIRECT_POLICY": 100000,

	//设置代理函数
	"OPT_PROXY_FUNC": 100001,
	"OPT_DEBUG":      100002,

	//指定用于tls.Client的TLS配置信息
	"OPT_UNSAFE_TLS": 100004,
}

// Default options for any clients.
var defaultOptions = map[int]interface{}{
	OPT_FOLLOWLOCATION: true,
	OPT_MAXREDIRS:      10,
	OPT_AUTOREFERER:    true,
	OPT_USERAGENT:      USERAGENT,
	OPT_COOKIEJAR:      true,
	OPT_DEBUG:          false,
}

// These options affect transport, transport may not be reused if you change any
// of these options during a request.
//这些选项会影响传输，如果在请求期间更改这些选项中的任何一个，传输可能不会被重用。
var transportOptions = []int{
	OPT_CONNECTTIMEOUT,
	OPT_CONNECTTIMEOUT_MS,
	OPT_PROXYTYPE,
	OPT_TIMEOUT,
	OPT_TIMEOUT_MS,
	OPT_INTERFACE,
	OPT_PROXY,
	OPT_PROXY_FUNC,
	OPT_UNSAFE_TLS,
}

// These options affect cookie jar, jar may not be reused if you change any of
// these options during a request.
//  jar，如果在请求期间更改这些选项中的任何一个，则jar可能不会被重用。
var jarOptions = []int{
	OPT_COOKIEJAR,
}

// Thin wrapper of http.Response(can also be used as http.Response).
// 响应体
type Response struct {
	*http.Response
}

// Read response body into a byte slice.
func (this *Response) ReadAll() ([]byte, error) {
	var reader io.ReadCloser
	var err error
	switch this.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(this.Body)
		if err != nil {
			return nil, err
		}
	default:
		reader = this.Body
	}

	defer reader.Close()
	return ioutil.ReadAll(reader)
}

// Read response body into string.
// 响应体转换为stirng
func (this *Response) ToString() (string, error) {
	bytes, err := this.ReadAll()
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// Powerful and easy to use HTTP client.
type HttpClient struct {
	// Default options of this client.
	// 默认选项
	options map[int]interface{}

	// Default headers of this client.
	// 默认头部信息
	Headers map[string]string

	// Options of current request.
	// 当前请求信息
	oneTimeOptions map[int]interface{}

	// Headers of current request.
	// 当前头部请求信息
	oneTimeHeaders map[string]string

	// Cookies of current request.
	// 当前cookie信息
	oneTimeCookies []*http.Cookie

	// Global transport of this client, might be shared between different
	// requests.
	transport http.RoundTripper

	// Global cookie jar of this client, might be shared between different
	// requests.
	jar http.CookieJar

	// Whether current request should reuse the transport or not.
	//当前是否可以重复使用Transport
	reuseTransport bool

	// Whether current request should reuse the cookie jar or not.
	//当前是否可以重复使用cookie
	reuseJar bool

	// Make requests of one client concurrent safe.【并发安全】
	lock *sync.Mutex

	withLock bool
}

// Set default options and headers.
func (this *HttpClient) Defaults(defaults Map) *HttpClient {
	options, headers := parseMap(defaults)

	// merge options
	if this.options == nil {
		this.options = options
	} else {
		for k, v := range options {
			this.options[k] = v
		}
	}

	// merge headers
	if this.Headers == nil {
		this.Headers = headers
	} else {
		for k, v := range headers {
			this.Headers[k] = v
		}
	}

	return this
}

// Begin marks the begining of a request, it's necessary for concurrent
// 对于并发 开始标记请求的开始
func (this *HttpClient) Begin() *HttpClient {
	this.lock.Lock()
	this.withLock = true

	return this
}

// Reset the client state so that other requests can begin.
func (this *HttpClient) reset() {
	this.oneTimeOptions = nil
	this.oneTimeHeaders = nil
	this.oneTimeCookies = nil
	this.reuseTransport = true
	this.reuseJar = true

	// nil means the Begin has not been called, asume requests are not
	// concurrent.
	if this.withLock {
		this.lock.Unlock()
	}
}

// Temporarily specify an option of the current request.
// 暂时指定当前请求的选项。
func (this *HttpClient) WithOption(k int, v interface{}) *HttpClient {
	if this.oneTimeOptions == nil {
		this.oneTimeOptions = make(map[int]interface{})
	}
	this.oneTimeOptions[k] = v

	// Conditions we cann't reuse the transport.
	//判断Transport是否可以重用
	if hasOption(k, transportOptions) {
		this.reuseTransport = false
	}

	// Conditions we cann't reuse the cookie jar.
	//不能重复使用cookie
	if hasOption(k, jarOptions) {
		this.reuseJar = false
	}

	return this
}

// Temporarily specify multiple options of the current request.
// 批量指定当前请求的选项。
func (this *HttpClient) WithOptions(m Map) *HttpClient {
	options, _ := parseMap(m)
	for k, v := range options {
		this.WithOption(k, v)
	}

	return this
}

// Temporarily specify a header of the current request.
// 临时指定当前请求的头
func (this *HttpClient) WithHeader(k string, v string) *HttpClient {
	if this.oneTimeHeaders == nil {
		this.oneTimeHeaders = make(map[string]string)
	}
	this.oneTimeHeaders[k] = v

	return this
}

// Temporarily specify multiple headers of the current request.
// 批量指定当前请求的头
func (this *HttpClient) WithHeaders(m map[string]string) *HttpClient {
	for k, v := range m {
		this.WithHeader(k, v)
	}

	return this
}

// Specify cookies of the current request.
// 指定当前请求的Cookie
func (this *HttpClient) WithCookie(cookies ...*http.Cookie) *HttpClient {
	this.oneTimeCookies = append(this.oneTimeCookies, cookies...)

	return this
}

// Start a request, and get the response.
// Usually we just need the Get and Post method.
func (this *HttpClient) Do(method string, url string, headers map[string]string,
	body io.Reader) (*Response, error) {

	//合并参数选项
	options := mergeOptions(defaultOptions, this.options, this.oneTimeOptions)

	// 合并头信息
	headers = mergeHeaders(this.Headers, headers, this.oneTimeHeaders)
	cookies := this.oneTimeCookies

	var transport http.RoundTripper
	var jar http.CookieJar
	var err error

	// transport
	if this.transport == nil || !this.reuseTransport {
		transport, err = prepareTransport(options)
		if err != nil {
			this.reset()
			return nil, err
		}
		//如果可以重复使用
		if this.reuseTransport {
			this.transport = transport
		}
	} else {
		transport = this.transport
	}

	// Jar指定cookie管理器
	// 如果Jar为nil，请求中不会发送cookie，回复中的cookie会被忽略
	if this.jar == nil || !this.reuseJar {
		jar, err = prepareJar(options)
		if err != nil {
			this.reset()
			return nil, err
		}

		if this.reuseJar {
			this.jar = jar
		}
	} else {
		jar = this.jar
	}

	// release lock[释放锁]
	this.reset()

	//指定处理重定向的策略
	redirect, err := prepareRedirect(options)
	if err != nil {
		return nil, err
	}

	//HTTP 客户端
	c := &http.Client{
		Transport:     transport,
		CheckRedirect: redirect,
		Jar:           jar,
	}

	req, err := prepareRequest(method, url, headers, body, options)
	if err != nil {
		return nil, err
	}

	//调试使用
	if debugEnabled, ok := options[OPT_DEBUG]; ok {
		if debugEnabled.(bool) {
			//DumpRequestOut类似DumpRequest，但会包括标准http.Transport类型添加的头域，如User-Agent
			dump, err := httputil.DumpRequestOut(req, true)
			if err == nil {
				fmt.Printf("%s\n", dump)
			}
		}
	}

	if jar != nil {
		jar.SetCookies(req.URL, cookies)
	} else {
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}
	}

	res, err := c.Do(req)

	return &Response{res}, err
}
