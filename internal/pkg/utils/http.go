package utils

type CommonResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// // GwHttpPost 发送POST请求
// func GwRawHttpPost(ctx context.Context, domain, uriStr string, body interface{}, language string) ([]byte, error) {
// 	link := fmt.Sprintf("%s%s", domain, uriStr)
// 	header := getHttpHeader(ctx, language)
// 	_, valBytes, err := HTTPPost(ctx, link, header, body)

// 	if err != nil {
// 		// ulog.Warn(ctx, err)
// 		return nil, err
// 	}

// 	return valBytes, nil
// }

// // GwRawHttpGet
// func GwRawHttpGet(ctx context.Context, domain, uriStr string, language string) ([]byte, error) {
// 	link := fmt.Sprintf("%s%s", domain, uriStr)
// 	header := getHttpHeader(ctx, language)
// 	valBytes, err := HTTPGet(ctx, link, header, nil)

// 	if err != nil {
// 		ulog.Warn(ctx, err)
// 		return nil, err
// 	}

// 	return valBytes, nil
// }
