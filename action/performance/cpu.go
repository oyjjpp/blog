package performance

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var Count int64 = 0

func CPUData(ctx *gin.Context) {
	fmt.Println(ctx.Request.URL)
	fib := 0
	for i := 0; i < 5000; i++ {
		fib = Fib(20)
		fmt.Println("fib = ", fib)
	}
	str := RandomStr(RandomInt(100, 500))
	str = fmt.Sprintf("Fib = %d; String = %s", fib, str)
	ctx.JSON(http.StatusOK, str)
}

func CPUtest(ctx *gin.Context) {
	var fib int
	index := Count
	arr := make([]int, index)
	var i int64
	for ; i < index; i++ {
		fib = Fib(20)
		arr[i] = fib
		fmt.Println("fib = ", fib)
	}
	time.Sleep(time.Millisecond * 500)
	str := fmt.Sprintf("Fib = %v", arr)
	ctx.JSON(http.StatusOK, str)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandomStr(num int) string {
	seed := time.Now().UnixNano()
	if seed <= 0 {
		seed = time.Now().UnixNano()
	}
	rand.Seed(seed)
	b := make([]rune, num)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func calCount() {
	timeInterval := time.Tick(time.Second)
	for {
		select {
		case i := <-timeInterval:
			Count = int64(i.Second())
		}
	}
}
