package queue

import (
	"github.com/gin-gonic/gin"
	"github.com/oyjjpp/blog/service/queue"
	"net/http"
)

type ProductReq struct {
	Message string `form:"message" json:"message"`
}

func Product(ctx *gin.Context) {
	res := map[string]interface{}{
		"code": 0,
		"msg":  "success",
	}

	form := ProductReq{}
	if err := ctx.ShouldBind(&form); err != nil {
		res["code"] = 19000
		res["msg"] = "参数异常"
		ctx.JSON(http.StatusOK, res)
		return
	}

	if form.Message == "" {
		res["code"] = 19001
		res["msg"] = "message参数异常"
		ctx.JSON(http.StatusOK, res)
		return
	}

	// 生产数据
	queue.SendMessage("topic-study", form.Message)

	ctx.JSON(http.StatusOK, res)
}
