package controllers

import (
	"github.com/astaxie/beego"
	"oyjblog/util"
)

// Operations about test
type TestController struct {
	beego.Controller
}

func (t *TestController) TypeAssertion()  {
	json := make(map[string]interface{})

	num := 15.5

	a, err := util.Int(num)
	if err != nil {
		t.ServeJSON()	
	}
	json["num"] = a
	//type结构只能使用在switch语句上
	//json["numType"] = num.(type)

	
	t.Data["json"] = json
	t.ServeJSON()
}