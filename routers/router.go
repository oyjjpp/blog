// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"blog/controllers"
	"blog/controllers/admin"
	"blog/controllers/algorithm"
	"blog/controllers/network"

	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/context"
)

func init() {
	/*
			//URL + 闭包函数组成
			//此context不是golang标准包的 "github.com/astaxie/beego/context"
			beego.Get("get_func", func(ctx *context.Context){
				ctx.Output.Body([]byte("hello beego"))
			})

		    //#####固定路由
		    beego.Router("/", &controllers.UserController{}, "*:GetAll")

			//包分目录情况
		    beego.Router("/admin/getall", &admin.UserController{}, "*:GetAll")

			//RESTful 请求模式
			//http://localhost:8080/admin/object/hjkhsbnmn123
		    beego.RESTRouter("/admin/object", &admin.ObjectController{})


		    //#####正则匹配

		    //自动匹配
		    //函数名小写 http://47.98.161.8:8080/object/getall
		    beego.AutoRouter(&controllers.ObjectController{})


			//自定义方法
			beego.Router("/get/user", &controllers.UserController{}, "get:Get")

	*/
	//命名空间
	/*
		ns := beego.NewNamespace("/api",
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
			beego.NSNamespace("/admin/object",
				beego.NSInclude(
					&admin.ObjectController{},
				),
			),
			beego.NSNamespace("/admin/user",
				beego.NSInclude(
					&admin.UserController{},
				),
			),
		)
	*/
	ns := beego.NewNamespace("api")
	registerURL(ns)
	beego.AddNamespace(ns)
}

func registerURL(ns *beego.Namespace) {
	ns.Router("/object/getlist", &controllers.ObjectController{}, "get:GetAll")
	ns.Router("/user/get", &controllers.UserController{}, "get:GetAll")
	ns.Router("/admin/object/getlist", &admin.ObjectController{}, "get:GetAll")
	ns.Router("/admin/user/get", &admin.UserController{}, "get:GetAll")

	//测试
	ns.Router("/test/type", &controllers.TestController{}, "*:TypeAssertion")
	ns.Router("/test/ip", &controllers.TestController{}, "*:IpChange")
	ns.Router("/test/typechange", &controllers.TestController{}, "*:TypeChange")
	ns.Router("/test/redis", &controllers.TestController{}, "*:Redis")

	//排序算法
	ns.Router("/algorithm/insertsort", &algorithm.SortController{}, "*:InsertSort")
	ns.Router("/algorithm/shellsort", &algorithm.SortController{}, "*:ShellSort")

	//加密算法
	ns.Router("/algorithm/aes", &algorithm.EncryptionController{}, "*:AesEncry")
	ns.Router("/algorithm/des", &algorithm.EncryptionController{}, "*:DesEncry")
	ns.Router("/algorithm/rsa", &algorithm.EncryptionController{}, "*:RsaEncry")
	ns.Router("/algorithm/rsasign", &algorithm.EncryptionController{}, "*:RsaSign")

	//网络编程
	ns.Router("/network/get", &network.HttpController{}, "*:Get")
	ns.Router("/network/getcontent", &network.HttpController{}, "*:GetContent")

	//输入输出

}
