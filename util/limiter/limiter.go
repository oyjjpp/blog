package limiter

import "context"

type GLimit struct {
	Num int
	C   chan struct{}
	Ctx context.Context
}

func NewGLimit(num int, Ctx context.Context) *GLimit {
	return &GLimit{
		Num: num,
		C:   make(chan struct{}, num),
		Ctx: Ctx,
	}
}

func (g *GLimit) Run(f func(ctx context.Context)) {
	g.C <- struct{}{}
	go func() {
		f(g.Ctx)
		<-g.C
	}()
}
