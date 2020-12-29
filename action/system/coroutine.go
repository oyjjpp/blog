// 协程
package system

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// GetCoroutineID
// 协程ID
func GetCoroutineID() int {
	var buf [64]byte
	// 调用其的go程的调用栈踪迹格式化后写入到buf中并返回写入的字节数
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
