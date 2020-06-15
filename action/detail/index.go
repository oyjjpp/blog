package detail

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	body := map[string]interface{}{
		"id":        ctx.PostForm("id"),
		"user_name": ctx.PostForm("usr"),
	}
	res := map[string]interface{}{
		"code": 0,
		"msg":  "success",
		"body": body,
	}
	ctx.JSON(http.StatusOK, res)
}
