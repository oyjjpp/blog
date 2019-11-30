package network

import (
	"github.com/astaxie/beego"
	"github.com/ddliu/go-httpclient"
)

// Operations about object
type HttpController struct {
	beego.Controller
}

//简单发送请求
func (o *HttpController) Get() {
	// data := make(map[string]interface{})
	// url := "http://ah2.zhangyue.com/zybk/api/channel/index?key=ch_free"
	// data["source"] = url
	// rs, _ := service.HttpGet(url, map[string]string{})
	// result := make(map[string]interface{})
	// err := json.Unmarshal([]byte(rs), &result)
	// if err != nil {

	// }
	// data["rs"] = result
	// o.Data["json"] = data
	// o.ServeJSON()
}

func (o *HttpController) GetContent() {
	data := make(map[string]interface{})
	url := "http://ah2.zhangyue.com/zybk/api/channel/index?key=ch_free"
	data["source"] = url

	httpclient.Defaults(httpclient.Map{
		httpclient.OPT_USERAGENT: "my awsome httpclient",
		"Accept-Language":        "en-us",
	})

	res, _ := httpclient.Get(url, map[string]string{
		"q": "news",
	})

	data["rs"] = res
	o.Data["json"] = data
	o.ServeJSON()
}
