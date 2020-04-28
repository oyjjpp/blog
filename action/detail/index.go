package detail

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	res := map[string]interface{}{
		"code": 0,
		"msg":  "success",
		"body": "detai_index",
	}
	ctx.JSON(http.StatusOK, res)
}
