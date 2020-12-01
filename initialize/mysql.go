package initialize

import (
	"log"
	"os"
	"time"

	"github.com/oyjjpp/blog/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Mysql
// 连接数据库
// 确认日志记录条件？
func Mysql() {

	// 设置mysql 日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Millisecond, // 慢 SQL 阈值
			LogLevel:      logger.Info,      // Log level
			Colorful:      false,            // 禁用彩色打印
		},
	)

	// https://github.com/go-sql-driver/mysql#dsn-data-source-name
	// dsn := "web:web1234!@#$@tcp(47.98.161.8:8004)/qmPlus?charset=utf8&parseTime=True&loc=Local"
	dsn := "web:web1234!@#$@tcp(47.98.161.8:8004)/learn?charset=utf8&parseTime=True&loc=Local"
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger}); err != nil {
		// panic("failed to connect database")
		os.Exit(0)
	} else {
		global.MysqlDB = db
        
        // 设置表名为非负数形式
        // global.MysqlDB.SingularTable(true)
        
		// 设置连接池
		if sqlDB, err := global.MysqlDB.DB(); err == nil {
			// SetMaxIdleConns 设置空闲连接池中连接的最大数量
			sqlDB.SetMaxIdleConns(2)
			// SetMaxOpenConns 设置打开数据库连接的最大数量。
			sqlDB.SetMaxOpenConns(10)
			// SetConnMaxLifetime 设置了连接可复用的最大时间。
			sqlDB.SetConnMaxLifetime(time.Hour)
		}
	}
}
