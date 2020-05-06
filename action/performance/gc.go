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
func avg(items []time.Duration) time.Duration {
	var sum time.Duration
	for _, item := range items {
		sum += item
	}
	return time.Duration(int64(sum) / int64(len(items)))
}
