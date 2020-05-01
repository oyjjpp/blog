package performance

import "testing"

func BenchmarkStrAddBuffer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = stringAddBuffer()
	}
	b.StopTimer()
}

func BenchmarkStrAddBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = stringAddBuilder()
	}
	b.StopTimer()
}

/*

基准测试手段

// 运行基准测试
go test -bench=.


// 排除其他测试
go test -bench=. -run=none
go test -run=^&

// 指定CPU核数
go test -bench=. -run=none -cpu=1,2,4

// 指定时间 以达到足够的采集信息
go test -bench=. -run=none  -benchtime=10s


常见导致性能的问题
1、执行时间过长
2、内存占用过高【错误的用法】
3、意外阻塞【锁竞争、协程阻塞】

常见捕获手段
1、net/http/pprof ： 适用于Web环境，直接通过访问http请求获取数据
2、runtime/pprof ： 代码中埋点记录监控数据
3、运行基准测试阶段：测试阶段生成采样数据


通过基准测试进行分析案例
测试阶段采集监控数据
// 首先收集内存数据
go test -bench=. -run=none -memprofile mem.out
// 采集CPU消耗数据
go test -bench=. -run=none -cpuprofile cpu.out
// 采集阻塞信息
go test -bench=. -run=none -blockprofile block.out

通过sample_index 选择其他信息
常见内存相关
inuse_space - 已分配但尚未释放的内存数量
inuse_objects - 已分配但尚未释放的对象数量
alloc_space - 已分配的内存总量（不管是否已释放）
alloc_objects - 已分配的对象总量（不管是否已释放）


go tool pprof 工具

// 查看单向统计信息
go tool pprof -text -alloc_object -cum mem.out
top 用法
// 输出源码统计信息
list
// 进一步查看目标,将一个函数里面每个执行命令拆开来分析
peek

*/
