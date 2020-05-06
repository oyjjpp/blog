// performance
// 性能分析
// 灵活性，适用于特定代码段的分析
// 两种方式
// 1、使用runtime/pprof
// 2、net/http/pprof
package performance

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)


func stringAdd() (str string) {
	for i := 0; i < 20000; i++ {
		str += strconv.Itoa(i)
	}
	return
}

func stringAddV2() (str string) {
	for i := 0; i < 20000; i++ {
		str = fmt.Sprintf("%s%d", str, i)
	}
	return
}

func stringAddBuffer() string {
	var str bytes.Buffer
	for i := 0; i < 20000; i++ {
		str.WriteString(strconv.Itoa(i))
	}
	return str.String()
}

func stringAddBuilder() string {
	var str strings.Builder
	for i := 0; i < 20000; i++ {
		str.WriteString(strconv.Itoa(i))
	}
	return str.String()
}
