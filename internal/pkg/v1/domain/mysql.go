package domain

import (
	"net/url"
	"strconv"
)

// CommonPage 分页
type CommonPage struct {
	Page  int64 `json:"page"`  // 第几页
	Size  int64 `json:"size"`  // 每页数
	Total int64 `json:"total"` // 总数
}

// GetCommonPage 获取分页参数
func GetCommonPage(query url.Values) *CommonPage {
	page, size := 1, 20
	if temp, err := strconv.Atoi(query.Get("page")); err == nil && temp > 0 {
		page = temp
	}
	if temp, err := strconv.Atoi(query.Get("size")); err == nil && temp > 0 {
		size = temp
	}

	return &CommonPage{
		Page: int64(page),
		Size: int64(size),
	}
}

// 表
const (
	TableUser = "user" //
)
