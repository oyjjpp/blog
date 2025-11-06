package domain

import "context"

type Limit struct {
	n   int
	c   chan struct{}
	ctx context.Context
}

func NewLimit(ctx context.Context, n int) *Limit {
	return &Limit{
		n:   n,
		c:   make(chan struct{}, n),
		ctx: ctx,
	}
}

// Run f in a new goroutine but with limit.
func (g *Limit) Run(f func(ctx context.Context)) {
	g.c <- struct{}{}
	go func() {
		f(g.ctx)
		<-g.c
	}()
}
