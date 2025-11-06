package config

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//go:embed config.yaml
var configYAML string

func ConfigInit() {
	viper.SetConfigType("yaml") // 配置文件类型

	if err := viper.ReadConfig(bytes.NewBufferString(configYAML)); err != nil {
		panic(fmt.Errorf("读取配置失败: %w", err))
	}

	// 3. 读取配置项（支持嵌套、默认值）
	// appName := viper.GetString("app.name") // 直接读取嵌套字段
	// appPort := viper.GetInt("app.port")    // 自动类型转换
	// dbHost := viper.GetString("mysql.host")
	// dbPass := viper.GetString("mysql.password")
	// dbPort := viper.GetInt("mysql.port")
	// fmt.Printf("app: %s:%d\n", appName, appPort)
	// fmt.Printf("db: %s:%d, user: %s, pass: %s\n", dbHost, dbPort, viper.GetString("mysql.user"), dbPass)

	// 4. 监听配置文件变更（动态更新）
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("配置文件变更: %s\n", e.Name)
	})
}
