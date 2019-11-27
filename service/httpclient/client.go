//wiki https://github.com/ddliu/go-httpclient/blob/master/httpclient.go
package httpclient

import (
	"compress/gzip"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

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
	OPT_UNSAFE_TLS      = 100004

	//长连接TCP池大小
	TcpMaxConnsPoll = 100
	//keepAlive的超时时间
	TcpKeepAliveTimeOut = 30 * time.Second
)

//全局公用的transport方便于做长连接和连接池
var gtransport http.RoundTripper

// Thin wrapper of http.Response(can also be used as http.Response).
type Response struct {
	*http.Response
}

// Read response body into string.
func (this *Response) ToString() (string, error) {
	bytes, err := this.ReadAll()
	if err != nil {
		return "", err
	}

	return string(bytes), nil
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

//http请求客户端 实现net/http 下client结构
type HttpClient struct {
	timeOut   time.Duration
	transport http.RoundTripper
}

//初始化
func init() {
	options := make(map[int]interface{})

	//设置代理
	options[OPT_PROXY_FUNC] = func(req *http.Request) (*url.URL, error) {
		proxyURL := req.Header.Get("proxyUrl")
		req.Header.Del("proxyUrl")
		if proxyURL == "" {
			return nil, nil
		}
		u, err := url.Parse(proxyURL)
		if err != nil {
			return nil, err
		}
		return u, nil
	}

	var err error
	gtransport, err = prepareTransport(options)
	if err != nil {
		panic(err)
	}
}

//构造函数
func NewClient() (*HttpClient, error) {
	return &HttpClient{transport: gtransport}, nil
}

//获取transport对象 Transport类型可以缓存连接以在未来重用
// Handles timemout, proxy and maybe other transport related options here.
func prepareTransport(options map[int]interface{}) (http.RoundTripper, error) {

	//连接超时时间设置
	connectTimeoutMS := 0
	if connectTimeoutMS_, ok := options[OPT_CONNECTTIMEOUT_MS]; ok {
		//类型断言 校验超时设置是否为int类型 设置单位为毫秒
		if connectTimeoutMS, ok = connectTimeoutMS_.(int); !ok {
			return nil, errors.New("OPT_CONNECTTIMEOUT_MS must be int")
		}
	} else if connectTimeout_, ok := options[OPT_CONNECTTIMEOUT]; ok {
		//设置单位为秒
		if connectTimeout, ok := connectTimeout_.(int); ok {
			connectTimeoutMS = connectTimeout * 1000
		} else {
			return nil, errors.New("OPT_CONNECTTIMEOUT must be int")
		}
	}

	//超时设置
	timeoutMS := 0
	if timeoutMS_, ok := options[OPT_TIMEOUT_MS]; ok {
		if timeoutMS, ok = timeoutMS_.(int); !ok {
			return nil, errors.New("OPT_TIMEOUT_MS must be int")
		}
	} else if timeout_, ok := options[OPT_TIMEOUT]; ok {
		if timeout, ok := timeout_.(int); ok {
			timeoutMS = timeout * 1000
		} else {
			return nil, errors.New("OPT_TIMEOUT must be int")
		}
	}

	//未设置连接超时时间或者连接超时时间大于http 超时时间则进行重置
	// fix connect timeout(important, or it might cause a long time wait during connection)
	if timeoutMS > 0 && (connectTimeoutMS > timeoutMS || connectTimeoutMS == 0) {
		connectTimeoutMS = timeoutMS
	}

	//实例化一个Transport结构
	transport := &http.Transport{
		//最大的限制连接
		MaxIdleConnsPerHost: TcpMaxConnsPoll,
	}

	//指定创建TCP连接的拨号函数，如果Dial为nil，则会使用net.Dial
	transport.Dial = func(network, addr string) (net.Conn, error) {
		var conn net.Conn
		var err error

		//建立连接的参数
		d := net.Dialer{
			//指定一个活动连接的声明周期
			KeepAlive: TcpKeepAliveTimeOut,
		}
		if connectTimeoutMS > 0 {
			//等待连接建立的最大时长（毫秒）
			d.Timeout = time.Duration(connectTimeoutMS) * time.Millisecond
		}

		conn, err = d.Dial(network, addr)
		if err != nil {
			return nil, err
		}

		//设置读写操作绝对期限
		if timeoutMS > 0 {
			conn.SetDeadline(time.Now().Add(time.Duration(timeoutMS) * time.Millisecond))
		}

		return conn, nil
	}

	// proxy 指定一个对给定请求返回代理的函数
	if proxyFunc_, ok := options[OPT_PROXY_FUNC]; ok {
		//通过函数设置代理
		if proxyFunc, ok := proxyFunc_.(func(*http.Request) (*url.URL, error)); ok {
			transport.Proxy = proxyFunc
		} else {
			return nil, errors.New("OPT_PROXY_FUNC is not a desired function")
		}
	} else {
		//设置类型及值设置代理
		var proxytype int
		//校验代理类型
		if proxytype_, ok := options[OPT_PROXYTYPE]; ok {
			if proxytype, ok = proxytype_.(int); !ok || proxytype != PROXY_HTTP {
				return nil, errors.New("OPT_PROXYTYPE must be int, and only PROXY_HTTP is currently supported")
			}
		}

		var proxy string
		if proxy_, ok := options[OPT_PROXY]; ok {
			if proxy, ok = proxy_.(string); !ok {
				return nil, errors.New("OPT_PROXY must be string")
			}
			proxyURL, err := url.Parse(proxy)
			if err != nil {
				return nil, err
			}
			transport.Proxy = http.ProxyURL(proxyURL)
		}
	}

	return transport, nil
}

//设置超时
func (this *HttpClient) SetTimeOut(timeOut time.Duration) {
	this.timeOut = timeOut
}

//GET 请求
func (this *HttpClient) Get(url string, params map[string]string, headers map[string]string) (*Response, error) {
	url = addParams(url, params)
	return this.Do("GET", url, headers, nil)
}

// Add params to a url string.
func addParams(url_ string, params map[string]string) string {
	if len(params) == 0 {
		return url_
	}

	if !strings.Contains(url_, "?") {
		url_ += "?"
	}

	if strings.HasSuffix(url_, "?") || strings.HasSuffix(url_, "&") {
		url_ += paramsToString(params)
	} else {
		url_ += "&" + paramsToString(params)
	}

	return url_
}

// Convert string map to url component.
func paramsToString(params map[string]string) string {
	values := url.Values{}
	for k, v := range params {
		values.Set(k, v)
	}

	return values.Encode()
}

//实际http请求函数
func (this *HttpClient) Do(method string, url string, headers map[string]string, body io.Reader) (*Response, error) {

	var err error

	// transport
	if this.transport == nil {
		return nil, errors.New("transport is nil")
	}

	//如果未设置超时时间则默认设置为1s
	if this.timeOut == time.Duration(0) {
		this.timeOut = time.Second * 1 //默认超时1s
	}

	//初始化client结构体
	c := &http.Client{
		Transport: this.transport,
		Timeout:   this.timeOut,
	}

	//准备请求
	req, err := this.prepareRequest(method, url, headers, body, map[int]interface{}{})
	if err != nil {
		return nil, err
	}

	res, err := c.Do(req)
	return &Response{res}, err
}

//获取request对象
func (this *HttpClient) prepareRequest(method string, url_ string, headers map[string]string, body io.Reader, options map[int]interface{}) (*http.Request, error) {

	//初始化一个新的请求体
	req, err := http.NewRequest(method, url_, body)

	if err != nil {
		return nil, err
	}

	// OPT_REFERER
	if referer, ok := options[OPT_REFERER]; ok {
		if refererStr, ok := referer.(string); ok {
			req.Header.Set("Referer", refererStr)
		}
	}

	// OPT_USERAGENT
	if useragent, ok := options[OPT_USERAGENT]; ok {
		if useragentStr, ok := useragent.(string); ok {
			req.Header.Set("User-Agent", useragentStr)
		}
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return req, nil
}
