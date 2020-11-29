package user

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/oyjjpp/blog/global"
	"github.com/oyjjpp/blog/models"
	"gorm.io/gorm"
)

func Read(ctx *gin.Context) {
	var user models.SysUser
	userName := ctx.Query("username")
	data := global.MysqlDB.Where("username = ?", userName).First(&user)
	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500204,
			"msg":  data.Error.Error(),
			"body": map[string]interface{}{},
		})
	} else if data.Error == nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500200,
			"msg":  "sucess",
			"body": user,
		})
	} else {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500200,
			"msg":  data.Error.Error(),
			"body": map[string]interface{}{},
		})
	}
}

// Table
// 指定Table
func Table(ctx *gin.Context) {
	result := map[string]interface{}{}
	data := global.MysqlDB.Table("sys_users").Take(&result)
	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500204,
			"msg":  data.Error.Error(),
			"body": result,
		})
	} else if data.Error == nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500200,
			"msg":  "sucess",
			"body": result,
		})
	} else {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500200,
			"msg":  data.Error.Error(),
			"body": result,
		})
	}
}

// SearchByKey
// 根据主键搜索
func SearchByKey(ctx *gin.Context) {
	id := ctx.Query("id")
	var result models.SysUser
	data := global.MysqlDB.First(&result, id)
	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500204,
			"msg":  data.Error.Error(),
			"body": result,
		})
	} else if data.Error == nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500200,
			"msg":  "sucess",
			"body": result,
		})
	} else {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500200,
			"msg":  data.Error.Error(),
			"body": result,
		})
	}
}

// SearchByKey
// 根据主键搜索多个
func SearchSByKey(ctx *gin.Context) {
	id := ctx.Query("id")
	var result []models.SysUser
	if id == "" {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500201,
			"msg":  "param error",
			"body": result,
		})
	}
	ids := strings.Split(id, ",")
	data := global.MysqlDB.Find(&result, ids)
	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500204,
			"msg":  data.Error.Error(),
			"body": result,
		})
	} else if data.Error == nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500200,
			"msg":  "sucess",
			"body": result,
		})
	} else {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500200,
			"msg":  data.Error.Error(),
			"body": result,
		})
	}
}

// SearchWhere
// 根据条件搜索
func SearchWhere(ctx *gin.Context) {
	id := ctx.Query("id")
	var result []models.SysUser
	if id == "" {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500201,
			"msg":  "param error",
			"body": result,
		})
	}
	ids := strings.Split(id, ",")
	data := global.MysqlDB.Find(&result, ids)
	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500204,
			"msg":  data.Error.Error(),
			"body": result,
		})
	} else if data.Error == nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500200,
			"msg":  "sucess",
			"body": result,
		})
	} else {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 500200,
			"msg":  data.Error.Error(),
			"body": result,
		})
	}
}
