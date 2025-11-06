package utils

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/valyala/fasthttp"
)

func FastRequestJson(ctx context.Context, method string, url string, header map[string]string, args interface{}) ([]byte, error) {
	req := &fasthttp.Request{}
	req.SetRequestURI(url)

	var body []byte
	var err error
	if newArge, ok := args.([]byte); ok {
		body = newArge
	} else {
		body, err = json.Marshal(args)
		if err != nil {
			// ulog.Debugf(ctx, "%v", err)
			return nil, err
		}
	}

	if strings.Contains(url, "pub/publish") {
		// ulog.Debugf(ctx, "link:%s header:%v", url, header)
	} else {
		// ulog.Debugf(ctx, "link:%s header:%v body:%v", url, header, util.BytesToString(body))
	}

	req.SetBody(body)
	req.Header.SetContentType("application/json")
	req.Header.SetMethod(method)
	if len(header) > 0 {
		for key, value := range header {
			req.Header.Add(key, value)
		}
	}

	resp := &fasthttp.Response{}
	if err := fasthttp.Do(req, resp); err != nil {
		// ulog.Debugf(ctx, "%v", err)
		return nil, err
	}
	if strings.Contains(url, "get_qrcode") {
		// ulog.Debugf(ctx, "link:%s", url)
	} else {
		// ulog.Debugf(ctx, "link:%s response:%v", url, util.BytesToString(resp.Body()))
	}

	return resp.Body(), nil
}
