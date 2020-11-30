package route

import (
	"github.com/gin-gonic/gin"
	"github.com/oyjjpp/blog/action/detail"
	"github.com/oyjjpp/blog/action/performance"
	"github.com/oyjjpp/blog/action/section"
	"github.com/oyjjpp/blog/action/serialize"
	"github.com/oyjjpp/blog/action/user"
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
	}
}
