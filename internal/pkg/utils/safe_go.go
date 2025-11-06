package utils

import (
	"context"
	"sync"
)

type groupGo struct {
	wg     *sync.WaitGroup
	HasErr bool
}

/*
示例：

	 var companyName string
	 var userName string

	 group := utils.NewGroupGo()
	 group.InitString(ctx, &companyName, func() (string, error) {
		// 耗时方法
		return c.getCompanyInfo(ctx, companyId)
	})

	group.InitString(ctx, &userName, func() (string, error) {
		// 耗时方法
		return c.getUserName(ctx, userId)
	})

	// 等待所有耗时方法执行结束
	group.Wait()

	if group.HasErr {
		ulog.Warn(ctx, "并发获取信息时出现了错误")
	}
*/
func NewGroupGo() *groupGo {
	return &groupGo{
		wg:     &sync.WaitGroup{},
		HasErr: false,
	}
}

func (c *groupGo) Wait() {
	c.wg.Wait()
}

// 并发执行 fn, 并将返回值string赋给filed指针
func (c *groupGo) InitString(ctx context.Context, field *string, fn func() (string, error)) {
	c.safeGo(ctx, func() {
		resp, err := fn()
		if err != nil {
			c.HasErr = true
			// ulog.Warn(ctx, err)
			return
		}
		*field = resp
	})
}

func (c *groupGo) InitAny(ctx context.Context, fn func() error) {
	c.safeGo(ctx, func() {
		err := fn()
		if err != nil {
			c.HasErr = true
			// ulog.Warn(ctx, err)
			return
		}
	})
}

func (c *groupGo) safeGo(_ context.Context, fn func()) {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		defer func() {
			if err := recover(); err != nil {
				c.HasErr = true
				// ulog.Warn(ctx, err)
			}
		}()

		fn()
	}()
}
