package utils

import (
	"context"

	"golang.org/x/time/rate"
)

var deployAlertChan = make(chan *MessageBody, 100)

func init() {
	robotLink := "https://open.feishu.cn/open-apis/bot/v2/hook/ea57ecec-c6cb-45e4-abfa-db0285fba73f"
	limiter := rate.NewLimiter(5, 5)
	go func() {
		for body := range deployAlertChan {
			if limiter.Allow() {
				_, _ = FastRequestJson(context.Background(), "POST", robotLink, map[string]string{}, body)
			}
		}
	}()
}

// 发送机器人消息
func SendRobotPostMessage(title string, errInfoList []interface{}) error {
	content := []interface{}{}
	for _, value := range errInfoList {
		item := []map[string]interface{}{
			{
				"tag":  "text",
				"text": value,
			},
		}
		content = append(content, item)
	}

	message := map[string]interface{}{
		"post": map[string]interface{}{
			"zh_cn": map[string]interface{}{
				"title":   title,
				"content": content,
			},
		},
	}

	body := &MessageBody{
		MsgType: "post",
		Content: message,
	}
	// robotLink := "https://open.feishu.cn/open-apis/bot/v2/hook/ea57ecec-c6cb-45e4-abfa-db0285fba73f"
	// FastRequestJson(context.Background(), "POST", robotLink, map[string]string{}, body)
	select {
	case deployAlertChan <- body:
	default:
		// xlog.Infof("报警消息队列积压过滤, 当前消息无法添加到队列, 消息内容为: title:%s", title)
	}

	return nil
}
