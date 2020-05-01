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

func FilePerformance() {
	// CPU性能分析
	fileCpu, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile", err)
	}
	if err := pprof.StartCPUProfile(fileCpu); err != nil {
		log.Fatal("could not start CPU profile", err)
	}
	defer pprof.StopCPUProfile()

	// CPU性能分析
	// runtime.GC()
	fileMem, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal("could not create memory profile", err)
	}
	if err := pprof.WriteHeapProfile(fileMem); err != nil {
		log.Fatal("could not start memory profile", err)
	}
	defer fileMem.Close()

	// goroutine
	fileGoroutine, err := os.Create("goroutine.prof")
	if err != nil {
		log.Fatal("could not create goroutine profile", err)
	}
	if gPprof := pprof.Lookup("goroutine"); gPprof != nil {
		log.Fatal("could not start goroutine profile", err)
	} else {
		gPprof.WriteTo(fileGoroutine, 0)
	}
	defer fileGoroutine.Close()
}

func Fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	default:
		return Fib(n-1) + Fib(n-2)
	}
}

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
