package network

import (
	"encoding/json"
	"oyjblog/service"

	"github.com/astaxie/beego"
)

// Operations about object
type HttpController struct {
	beego.Controller
}

//插入排序算法
func (o *HttpController) Get() {
	data := make(map[string]interface{})
	url := "http://ah2.zhangyue.com/zybk/api/detail/index?key=17B11870175"
	data["source"] = url
	rs := service.HttpGet(url)
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(rs), &result)
	if err != nil {

	}
	data["rs"] = result
	o.Data["json"] = data
	o.ServeJSON()
}
