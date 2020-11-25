package route

import (
	"github.com/gin-gonic/gin"
	"github.com/oyjjpp/blog/action/detail"
	"github.com/oyjjpp/blog/action/performance"
	"github.com/oyjjpp/blog/action/section"
	"github.com/oyjjpp/blog/action/serialize"
)

func LoadRoute(engine *gin.Engine) {
	// route
	demoRouter := engine.Group("/blog")
	{
		demoRouter.POST("/detail/list", detail.List)
		demoRouter.POST("/detail/index", detail.Index)
		demoRouter.GET("/section/index", section.Index)

		// performance
		demoRouter.GET("/performance/cpu_data", performance.CPUData)
		demoRouter.GET("/performance/cpu_test", performance.CPUtest)
		demoRouter.GET("/performance/gc_data", performance.GcData)

		// serialize
		demoRouter.GET("/serialize/protobuf", serialize.Protobuf)
	}
}
