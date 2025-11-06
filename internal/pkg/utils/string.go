package utils

import (
	"crypto/rand"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// RandString 随机生成指定长度的字串
func RandString(length ...int) string {
	l := 32
	if len(length) > 0 && length[0] > 0 {
		l = length[0]
	}
	b := make([]byte, l/2)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	return fmt.Sprintf("%x", b)
}

// 针对手机号进行脱敏
func EncryptPhone(phone string) string {
	if len(phone) < 7 {
		return phone
	}
	temp := []byte(phone)
	for i := 3; i < 8; i++ {
		temp[i] = '*'
	}
	return string(temp)
}

func FormatHeader(header map[string]string) string {
	return fmt.Sprintf("%+v", header)
}

// 日志中将body转换为一行的string
func FormatBodyToString(body []byte) string {
	if len(body) > 1024000 {
		return "body 大于1兆,不显示全部, " + strings.ReplaceAll(string(body[:1024000]), "\r\n", "")
	}
	return strings.ReplaceAll(string(body), "\r\n", "")
}

func ConvertToString[T int | int64](list []T, sep string) string {
	var strList = make([]string, 0, len(list))
	for _, v := range list {
		strList = append(strList, fmt.Sprintf("%d", v))
	}
	return strings.Join(strList, sep)
}

func ConvertToSlice[T int | int64](str string, sep string) []T {
	var list []T
	for _, v := range strings.Split(str, sep) {
		num, _ := strconv.ParseInt(v, 10, 64)
		if num < 1 {
			continue
		}
		list = append(list, T(num))
	}
	return list
}

func UniqueSlice[T int | int64](list []T) []T {
	var filter = make(map[T]bool)
	var newList = make([]T, 0, len(list))
	for _, v := range list {
		if filter[v] {
			continue
		}
		filter[v] = true
		newList = append(newList, v)
	}
	return newList
}

func Paginate[T any](pg int, size int, list []T) []T {
	if pg < 1 {
		return nil
	}

	// 计算需要跳过的列表项数量
	skip := (pg - 1) * size

	// 检查是否需要跳过的数量超过列表长度
	if skip >= len(list) {
		return nil
	}

	// 计算结束索引
	end := skip + size
	if end > len(list) {
		end = len(list)
	}

	// 提取当前页的数据
	currentPage := list[skip:end]

	return currentPage
}

func DataInSlice[T int | int64 | string | int32](s []T, item T) bool {
	for _, s := range s {
		if s == item {
			return true
		}
	}
	return false
}

// DataDifference 求两个切片的差集
func DataDifference[T int | int64](a, b []T) []T {
	var diffArray []T
	temp := map[T]struct{}{}

	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range a {
		if _, ok := temp[val]; !ok {
			diffArray = append(diffArray, val)
		}
	}

	return diffArray
}

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

// ClearNonAlphanumericRegex
func ClearNonAlphanumericRegex(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}
