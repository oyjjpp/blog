package system

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// with_value
// 传值
func with_value() {
	// 声明一个根上下文信息
	ctx, cancel := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx, "key", "add value")

	go watch(valueCtx)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func watch(ctx context.Context) {
	fmt.Println("go watch")
	for {
		select {
		case <-ctx.Done():
			// 上下文被取消
			fmt.Println(ctx.Value("key"), "is cancel")
			time.Sleep(1 * time.Second)
		default:
			fmt.Println(ctx.Value("key"), "int goroutine")
			time.Sleep(1 * time.Second)
		}
	}
}

// with_timeout
// 超时
func with_timeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	fmt.Println("with_timeout start")
	var wg sync.WaitGroup
	wg.Add(1)
	go work(ctx, &wg)
	wg.Wait()
	fmt.Println("with_timeout end")
}

func work(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("指定时间", i)
		case <-ctx.Done():
			fmt.Println("Cancel context", i)
			return ctx.Err()
		}
	}
	return nil
}

// with_deadtime
// 截止时间
func with_deadtime() {
	dt := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), dt)
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-time.After(2 * time.Second):
				fmt.Println("到达截止时间")
			case <-ctx.Done():
				fmt.Println("上下文取消")
				return
			}
		}
	}(ctx)

	time.Sleep(5 * time.Second)
}

// 三个任务消耗时长不一样
func task1() int {
	time.Sleep(3 * time.Second)
	return 1
}

func task2() int {
	time.Sleep(2 * time.Second)
	return 2
}

func task3() int {
	time.Sleep(2 * time.Second)
	return 3
}

// with_url
// 一个请求触发三个任务 并发完成，出现一个超时则全部取消
// 1、要求处理三个结果之和
// 2、返回时间不超过指定时间
func with_url() {
	var res, sum int
	// 请求标志
	sucess := make(chan int, 1)
	// 保存每个任务的结果
	resChan := make(chan int, 3)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 使用sync保证任务全部完成之后推出
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		resChan <- task1()
		wg.Done()
	}()
	go func() {
		resChan <- task2()
		wg.Done()
	}()
	go func() {
		resChan <- task3()
		wg.Done()
	}()

	go func() {
		for {
			select {
			case res = <-resChan:
				sum += res
				fmt.Println("add", res)
			case <-sucess:
				fmt.Println("所有任务完成之后的结果", sum)
				// wg.Done()
				return
			case <-ctx.Done():
				// close(sucess)
				fmt.Println("出现超时后的结果", sum, ctx.Err().Error())
				// wg.Done()
				return
			}
		}
	}()

	wg.Wait()
	sucess <- 1
	// 表示任务已经完成 并且没有出现超时
	// if rs := ctx.Err(); rs == nil {
	// 	fmt.Println("ctx status", ctx.Err())
	// 	sucess <- 1
	// }
	return
}
