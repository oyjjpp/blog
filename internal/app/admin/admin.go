package admin

import (
	"api/config"
	_ "api/internal/pkg/lang" // 导入语言库
	"api/internal/pkg/store"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {
	// 初始化配置
	config.ConfigInit()
	// lang.Init()
	// // 日志初始化
	// domain.NewProductionLogger()
	// defer domain.Logger.Sync()

	// 存储初始化
	err := store.Init()
	if err != nil {
		panic(err)
	}

	// id生成器初始化
	// domain.InitIdGen(1)

	// 创建 Gin 引擎
	// gin.SetMode(gin.ReleaseMode)
	// route := gin.Default()
	route := gin.New()
	route.Use(gin.Recovery())

	// 启动后台进程
	// ctx := context.Background()

	// fmt.Printf("service name uuid: %s\n", adminUUID)

	// // 直播相关
	// go RunLiveCron(ctx)

	// // 运营推送相关
	// go RunPushCron(ctx)

	// // gorse事件相关
	// go RunGorseEventCron(ctx)

	// // 手机型号相关
	// go RunMobileModelCron(ctx)

	// // 公告同步相关
	// go RunAnnouncementSyncCron(ctx)

	// 注册路由
	// userOrder := usecase.NewUserOrderInfrastructure(mysql.MysqlDB, redis.RedisClient, config.GetData().Kafka)

	// handler := admin.NewAdminApiHandler(ctx, userOrder)
	// handler.RegisterRoutes(route)

	// 启动HTTP服务器
	// ulog.Infof(ctx, "管理后台服务器启动在端口 :9003")
	if err := route.Run(fmt.Sprintf(":%d", viper.GetInt("app.port"))); err != nil {
		// ulog.Errorf(ctx, "管理后台服务器启动失败: %v", err)
	}
}
