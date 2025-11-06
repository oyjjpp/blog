// Package ctime 用于自定义事件转换
package ctime

import (
	"time"
)

// GetUTCZeroTimeString 获取 utc 0 时区 时间 用于系统自己产生时间
func GetUTCZeroTimeString(layout string) string {
	return time.Now().UTC().Format(layout)
}

// FormatZeroTime2RFC3339 格式化为 0 时区的 UTC 时间  用于将存储的时间转换给前端 2006-01-02 15:05:05 => 2006-01-02T15:05:05Z
func FormatZeroTime2RFC3339(ts, layout string) string {
	if ts == "0000-00-00 00:00:00" {
		return ""
	}
	t, _ := time.ParseInLocation(layout, ts, time.Now().UTC().Location())
	return t.Format(time.RFC3339)
}

// ZeroRFC3339Time2Layout 前端传过来的 UTC 时间转换成数据库入库时间格式 2006-01-02T15:05:05Z => 2006-01-02 15:05:05
func ZeroRFC3339Time2Layout(ts, layout string) string {
	t, _ := time.ParseInLocation(time.RFC3339, ts, time.Now().UTC().Location())
	return t.Format(layout)
}

// ZeroTime2LocalTimeString 0时区 UTC 时间 转 本地时间 用于接收前端时间并转换为本地时间 //  用不上了
func ZeroTime2LocalTimeString(ts, layout string) string {
	t, _ := time.ParseInLocation(time.RFC3339, ts, time.UTC)
	return t.Local().Format(layout)
}
