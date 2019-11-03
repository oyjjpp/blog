package util

import "time"

//获取两个时间的时间间隔
func GetTimeSub(startTime, endTime time.Time) time.Duration {
	return endTime.Sub(startTime)
}

//获取当前到指定时间的时间间隔
func GetTimeSince(startTime time.Time) time.Duration {
	return time.Since(startTime)
}

//获取当前时间
func GetCurrentTime(layout string) string {
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}
	return time.Now().Format(layout)
}

//获取当前时间戳
func GetCurrentTimeUnix() int64 {
	return time.Now().Unix()
}

//将时间字符串转换成time类型，当前时区
func FormatStringToTime(layout, value string) (time.Time, error) {
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}
	return time.ParseInLocation(layout, value, time.Now().Location())
}

//将整型转换为time类型
func FormatIntToTime(value int64) time.Time {
	return time.Unix(value, 0)
}
