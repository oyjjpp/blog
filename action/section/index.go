package section

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	sync.Map
	res := map[string]interface{}{
		"code": 0,
		"msg":  "success",
		"body": "section_index",
	}
	ctx.JSON(http.StatusOK, res)
}
