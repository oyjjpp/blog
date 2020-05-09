// 垃圾回收数据
package performance

import (
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
)

var appStartTime = time.Now()

// GcData
func GcData(ctx *gin.Context) {
	gcstats := &debug.GCStats{PauseQuantiles: make([]time.Duration, 100)}
	debug.ReadGCStats(gcstats)
	var data = make(map[string]interface{})

	if gcstats.NumGC > 0 {
		elapsed := time.Now().Sub(appStartTime)
		overhead := float64(gcstats.PauseTotal) / float64(elapsed) * 100
		data["numGC"] = gcstats.NumGC
		data["pause"] = gcstats.Pause[0]
		data["pause_avg"] = avg(gcstats.Pause)
		data["overhead"] = overhead
	}
	ctx.JSON(http.StatusOK, data)
}

// avg
// 每次暂停收集垃圾的消耗的时间
func avg(items []time.Duration) time.Duration {
	var sum time.Duration
	for _, item := range items {
		sum += item
	}
	return time.Duration(int64(sum) / int64(len(items)))
}

/*

垃圾收集器的工作方式

收集器行为

三个阶段

标记设置-STW
标记-并发
标记终止-STW


收集开始时，会打开写屏障
写屏障的目的：允许收集器在收集期间保持堆上的数据完整性 因为收集器和应用程序是并发执行的

打开写屏障时，必须停止运行每个协程

GC会进行协程调度

GC 跟踪
GODEBUG=gctrace=1 ./app

减轻堆的压力，则将减少延迟成本，从而提高应用程序的性能

通过找到增加两个收集之间的时间的方案来降低收集开始的速度不是一个好策略
保持一致的速度会更好地保持应用程序以最佳性能运行

*/
