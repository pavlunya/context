// Package fluc provides fluent interface wrapper to work with context.
package fluc

import (
	"context"
	"time"
)

// chain is a wrapper for context.
type chain struct {
	ctx context.Context
}

// Context is a constructor, it returns new instance of chain.
func Context(ctx context.Context) *chain {
	return &chain{ctx}
}

// With wraps context.WithValue and returns chain.
// This method simplifies adding a lot of values to context:
//
// 	ctx := fluc.Context(ctx.Background()).
// 		With("key1", "value1").
// 		With("key2", "value2").
// 		Get()
//
// 	// In this case panic will never be caused because key1 is present in the context.
// 	// It's some kind of best practice to check if everything is ok.
// 	val, ok := ctx.Value("key1").(string)
// 	if !ok {
// 		panic("Something is absolutely broken!")
// 	}
//
// 	fmt.Printf("Some epic value is %s", val)
//
// Injecting values to the context was never so awesome.
func (c *chain) With(key interface{}, val interface{}) *chain {
	c.ctx = context.WithValue(c.ctx, key, val)
	return c
}

// WithDeadline wraps context.WithDeadline
func (c *chain) WithDeadline(deadline time.Time) (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithDeadline(c.ctx, deadline)
	c.ctx = ctx
	return
}

// WithCancel wraps context.WithCancel
func (c *chain) WithCancel() (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithCancel(c.ctx)
	c.ctx = ctx
	return
}

// WithTimeout wraps context.WithTimeout
func (c *chain) WithTimeout(timeout time.Duration) (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithTimeout(c.ctx, timeout)
	c.ctx = ctx
	return
}

// Get returns context. Well, as expected.
func (c *chain) Get() context.Context {
	return c.ctx
}
