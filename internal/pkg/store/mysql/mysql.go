package mysql

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DefaultDB *gorm.DB
)

func Init() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt64("mysql.port"),
		viper.GetString("mysql.dbname"),
		viper.GetString("mysql.charset"),
	)
	// newLogger := logger.New(
	// 	log.New(jsonFormatter{os.Stdout}, "", 0),
	// 	// log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出位置，可以是任何 io.Writer，标准输出是 os.Stdout）
	// 	logger.Config{
	// 		SlowThreshold: time.Second, // 慢 SQL 阈值
	// 		LogLevel:      logger.Info, // 日志等级
	// 		Colorful:      false,       // 彩色打印
	// 	},
	// )
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: nil,
	})
	if err != nil {
		zap.S().Fatalf("Failed to connect to database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		zap.S().Fatalf("Failed to get database: %v", err)
	}
	sqlDB.SetMaxIdleConns(viper.GetInt("mysql.max_idle"))
	sqlDB.SetMaxOpenConns(viper.GetInt("mysql.max_conn"))
	DefaultDB = db
	return nil
}
