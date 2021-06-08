package route

import (
	"github.com/gin-gonic/gin"
	"github.com/oyjjpp/blog/action/cache"
	"github.com/oyjjpp/blog/action/detail"
	"github.com/oyjjpp/blog/action/performance"
	"github.com/oyjjpp/blog/action/queue"
	"github.com/oyjjpp/blog/action/section"
	"github.com/oyjjpp/blog/action/serialize"
	"github.com/oyjjpp/blog/action/user"
	"net/http"
)

func LoadRoute(engine *gin.Engine) {
	// route
	demoRouter := engine.Group("/blog")
	{
		demoRouter.GET("/detail/list", detail.List)
		demoRouter.GET("/detail/index", detail.Index)
		demoRouter.GET("/section/index", section.Index)

		// performance
		demoRouter.GET("/performance/cpu_data", performance.CPUData)
		demoRouter.GET("/performance/cpu_test", performance.CPUtest)
		demoRouter.GET("/performance/gc_data", performance.GcData)

		// serialize
		demoRouter.GET("/serialize/protobuf", serialize.Protobuf)

		// mysql
		demoRouter.GET("/user/read", user.Read)
		demoRouter.GET("/user/table", user.Table)
		demoRouter.GET("/user/searchbykey", user.SearchByKey)
		demoRouter.GET("/user/searchsbykey", user.SearchSByKey)
		demoRouter.GET("/user/searchwhere", user.SearchWhere)
		demoRouter.GET("/user/batchcreate", user.BatchCreate)

		// cache
		demoRouter.GET("/cache/http", cache.Http)

		// http状态码
		demoRouter.GET("/api/error", blogError)
		demoRouter.GET("/api/redirect", blogRedirect)

		// mq
		demoRouter.GET("/kafka/product", queue.Product)

		demoRouter.POST("/kafka/:productID/appID", func(ctx *gin.Context) {
			productID := ctx.Param("productID")
			ctx.JSON(http.StatusOK, gin.H{
				"appId": productID,
			})
		})
	}
}

func blogError(ctx *gin.Context) {
	// ctx.Status(401)
	ctx.String(401, "%s", "请认证后访问")
}

func blogRedirect(ctx *gin.Context) {
	ctx.Redirect(301, "https://www.baidu.com")
}
