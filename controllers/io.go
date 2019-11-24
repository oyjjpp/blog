package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Users
type IoController struct {
	beego.Controller
}

func (u *IoController) Post() {
	uid := "input/output"
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}
