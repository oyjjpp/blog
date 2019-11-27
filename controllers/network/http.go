package network

import (
	"blog/service"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about object
type HttpController struct {
	beego.Controller
}

//简单发送请求
func (o *HttpController) Get() {
	data := make(map[string]interface{})
	url := "http://ah2.zhangyue.com/zybk/api/channel/index?key=ch_free"
	data["source"] = url
	rs, _ := service.HttpGet(url, map[string]string{})
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(rs), &result)
	if err != nil {

	}
	data["rs"] = result
	o.Data["json"] = data
	o.ServeJSON()
}
