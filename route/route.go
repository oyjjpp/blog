package route

import (
	"github.com/gin-gonic/gin"
	"github.com/oyjjpp/blog/action/detail"
	"github.com/oyjjpp/blog/action/performance"
	"github.com/oyjjpp/blog/action/section"
)

func LoadRoute(engine *gin.Engine) {
	// route
	demoRouter := engine.Group("/blog")
	{
		demoRouter.GET("/detail/index", detail.Index)
		demoRouter.GET("/section/index", section.Index)
		// performance
		demoRouter.GET("/performance/cpu_data", performance.CPUData)
		demoRouter.GET("/performance/cpu_test", performance.CPUtest)
	}
}
