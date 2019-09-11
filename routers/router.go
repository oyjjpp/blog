// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"oyjblog/controllers"
	"oyjblog/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
    //固定路由
    beego.Router("/", &controllers.UserController{}, "*:GetAll")
    //包分目录情况
    beego.Router("/admin/getall", &admin.UserController{}, "*:GetAll")


    //RESTful 请求模式
    beego.RESTRouter("/admin/object", &admin.ObjectController{})
    

    //正则匹配

    //自动匹配
    //函数名小写 http://47.98.161.8:8080/object/getall
    beego.AutoRouter(&controllers.ObjectController{})
    
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
