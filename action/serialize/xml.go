package serialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Xml(ctx *gin.Context) {
	res := map[string]interface{}{
		"serialize": "序列化",
	}
	ctx.JSON(http.StatusOK, res)
}
