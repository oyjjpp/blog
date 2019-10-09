package util

import (
	"bytes"
	"strconv"
	"strings"
)

//link https://www.cnblogs.com/hitfire/articles/6597654.html
//IpStringToInt IP 转换成 int类型
func IpStringToInt(ip string) int {
	//192.168.1.1 每8位使用一个字节保存
	ipSegs := strings.Split(ip, ".")
	var ipInt int = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}

//IpIntToString  IP 转换成string类型
func IpIntToString(ip int) string {
	ipSegs := make([]string, 4)
	var len int = len(ipSegs)
	buffer := bytes.NewBufferString("")

	for i := 0; i < len; i++ {
		tempInt := ip & 0xFF
		ipSegs[len-i-1] = strconv.Itoa(tempInt)
		ip = ip >> 8
	}

	for i := 0; i < len; i++ {
		buffer.WriteString(ipSegs[i])
		if i < len-1 {
			buffer.WriteString(".")
		}
	}
	return buffer.String()
}
