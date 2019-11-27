package controllers

import (
	"fmt"
	"blog/service"
	"blog/util"

	"github.com/astaxie/beego"
)

// Operations about test
type TestController struct {
	beego.Controller
}

func (t *TestController) TypeAssertion() {
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

//IP转换
func (t *TestController) IpChange() {
	ip := t.GetString("ip")
	rs := make(map[string]interface{})

	rs["ip"] = ip

	ipSave := util.IpStringToInt(ip)
	rs["ipToInt"] = ipSave
	rs["ipToString"] = util.IpIntToString(ipSave)

	t.Data["json"] = rs
	t.ServeJSON()
}

func (t *TestController) TypeChange() {
	rs := make(map[string]interface{})
	var number int = 7 / 2
	rs["num"] = 7.5 / 2
	rs["number"] = number
	t.Data["json"] = rs
	t.ServeJSON()
}

func (t *TestController) Redis() {
	rp := service.RedisPool{}
	redis_info_array := []string{"127.0.0.1:19000", "127.0.0.1:6399"}
	rp.InitRedisPool(redis_info_array)
	err := rp.RedisSet("learn", "golang-redis-pool")
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	rs, err := rp.RedisGet("learn")
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Println(rs)
}
