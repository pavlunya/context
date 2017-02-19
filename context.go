package fluc

import (
	"context"
	"time"
)

type chain struct {
	ctx context.Context
}

func Context(ctx context.Context) *chain {
	return &chain{ctx}
}

func (c *chain) With(key interface{}, val interface{}) *chain {
	c.ctx = context.WithValue(c.ctx, key, val)

	return c
}

func (c *chain) WithDeadline(deadline time.Time) (context.Context, context.CancelFunc) {
	return context.WithDeadline(c.ctx, deadline)
}

func (c *chain) WithCancel() (context.Context, context.CancelFunc) {
	return context.WithCancel(c.ctx)
}

func (c *chain) WithTimeout(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(c.ctx, timeout)
}

func (c *chain) Get() context.Context {
	return c.ctx
}