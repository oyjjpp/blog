package service

import (
	"blog/service/httpclient"
)

//HttpGet 对外http get请求接口
func HttpGet(path string, params map[string]string) (string, error) {
	//补充公用参数
	params = handleParam(params)

	args := HTTPArgs{
		Path:   path,
		Params: params,
	}
	gclient, err := httpclient.NewClient()
	if err != nil {
		return "", err
	}
	client := &HTTPClient{gclient}
	return client.Get(args)
}

//添加通用参数
func handleParam(params map[string]string) map[string]string {
	//设置默认超时时间
	if _, ok := params["timeout"]; !ok {
		params["timeout"] = "1"
	}

	//设置请求重试次数
	if _, ok := params["TryCount"]; !ok {
		params["TryCount"] = "2"
	}
	return params
}
