package initialize

import (
	"os"

	"github.com/oyjjpp/blog/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Mysql
func Mysql() {
	dsn := "web:web1234!@#$@tcp(47.98.161.8:8004)/qmPlus?charset=utf8&parseTime=True&loc=Local"
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		// panic("failed to connect database")
		os.Exit(0)
	} else {
		global.DB = db
	}
}
