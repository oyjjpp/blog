package detail

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	data := ``
	var info interface{}
	_ = json.Unmarshal([]byte(data), &info)
	ctx.JSON(http.StatusOK, info)
}
