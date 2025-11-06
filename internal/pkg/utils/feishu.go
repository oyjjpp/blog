package utils

type BaseRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type MessageBody struct {
	MsgType string                 `json:"msg_type"`
	Content map[string]interface{} `json:"content"`
}

type FeishuMessage struct {
	Tag  string `json:"tag"` // text, a
	Text string `json:"text"`
	Href string `json:"href,omitempty"`
}

// SendFeishuMessage 可以加 超链接的消息
// func SendFeishuMessage(ctx context.Context, link string, title string, message []*FeishuMessage) error {

// 	message = append(message, &FeishuMessage{
// 		Tag:  "text",
// 		Text: fmt.Sprintf("\napi_trace_id:%s\n", "ulog.GetApiTraceId(ctx)"),
// 	})

// 	messageMap := map[string]interface{}{
// 		"post": map[string]interface{}{
// 			"zh_cn": map[string]interface{}{
// 				"title":   fmt.Sprintf("(%s)[%s] %s", "config.Data().Mode", "operation-platform", title),
// 				"content": []interface{}{message},
// 			},
// 		},
// 	}

// 	body := MessageBody{
// 		MsgType: "post",
// 		Content: messageMap,
// 	}

// 	var header = http.Header{}
// 	header.Add("Content-Type", "application/json; charset=utf-8")

// 	_, _, err := HTTPPost(ctx, link, header, body)
// 	return err
// }
