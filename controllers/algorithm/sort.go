package algorithm

import (
	"fmt"
	"oyjblog/models"

	"github.com/astaxie/beego"
)

// Operations about object
type SortController struct {
	beego.Controller
}

//插入排序算法
func (o *SortController) InsertSort() {
	var dataItem = [7]int{5, 2, 0, 1, 3, 1, 4}
	data := make(map[string]interface{})
	data["source"] = dataItem
	fmt.Printf("add=%p, value=%v, type=%T\n", &dataItem, dataItem, dataItem)

	rs := models.InsertSort(dataItem[:])
	data["dest"] = rs
	fmt.Printf("add=%p, value=%v, type=%T\n", &rs, rs, rs)

	fmt.Printf("add=%p, value=%v, type=%T\n", &data, data, data)
	o.Data["json"] = data
	o.ServeJSON()
}

//希尔排序
func (o *SortController) ShellSort() {
	var dataItem = [7]int{5, 2, 0, 1, 3, 1, 4}

	data := make(map[string]interface{})
	data["source"] = dataItem
	rs := models.HeadSort(dataItem[:])
	data["dest"] = rs

	o.Data["json"] = data
	o.ServeJSON()
}
