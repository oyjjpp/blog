package initialize

import (
	"github.com/oyjjpp/blog/global"
	"github.com/oyjjpp/blog/models"
)

// DBTables
// 注册数据库表专用
func DBTables() {
	db := global.DB
	db.AutoMigrate(models.SysUser{})
}
