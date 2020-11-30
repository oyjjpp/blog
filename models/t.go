package models

type T struct {
	Id   int    `json:"id" gorm:"comment:'id'"`
	City string `json:"city"  gorm:"comment:'城市'"`
	Name string `json:"name" gorm:"default:'';comment:'用户昵称'" `
	Age  int    `json:"age" gorm:"default:0;comment:'年龄'"`
	Addr string `json:"addr" gorm:"default:'';comment:'用户角色ID'"`
}
