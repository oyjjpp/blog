package section

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	res := map[string]interface{}{
		"code": 0,
		"msg":  "success",
		"body": "section_index",
	}
	ctx.JSON(http.StatusOK, res)
}
