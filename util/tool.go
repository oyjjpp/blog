package util

import (
	"io"
	"os"
)

// writeFile 写入文件
func WriteFile(content string) {
	var err error
	var exist = true
	var logFile *os.File
	fileName := "./curl.log"
	// 先判断文件是否还存在
	if _, err = os.Stat(fileName); os.IsNotExist(err) {
		exist = false
	}
	if exist {
		//打开文件
		logFile, err = os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0666)
	} else {
		logFile, err = os.Create(fileName)
	}
	defer logFile.Close()
	if err != nil {
		return
	}
	_, err = io.WriteString(logFile, content+"\n") //写入文件(字符串)
}
