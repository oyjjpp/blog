package cache

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Http(ctx *gin.Context) {

	var data interface{}
	_ = json.Unmarshal([]byte(CACHE_DATA), &data)

	// 私有、共享
	// Cache-Control: private
	// Cache-Control: public

	// 缓存但验证、不缓存
	// Cache-Control: no-cache
	// Cache-Control: no-store

	// 缓存时长
	// Cache-Control: max-age=31536000
	// max-age > Expires > Last-Modified【(Date-Last-Modified)*10%】

	// 验证时长
	// Cache-Control: must-revalidate

	// Cache-Control:no-cache,no-store, must-revalidate
	ctx.Header("Cache-Control", "no-cache,max-age=604800")
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"msg":  "success",
		"body": data,
	})
}

// Squid
// (from disk cache)
// (from memory cache)
