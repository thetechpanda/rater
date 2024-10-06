package rater

import (
	"context"
	"sync/atomic"
	"time"
)

// Monitor is a simple rate monitor, it counts the rate of calls to Rate() method.
// Value() returns the current count.
// It resets the count every 'every' duration.
type Monitor struct {
	count atomic.Int64
	// C is a channel that will receive the current count every 'every' duration, then the count will be reset.
	C chan int64
}

// NewMonitor initialises and starts a monitor, once started a monitor can be stopped by cancelling the context passed in argument.
// If 'every' is less than or equal to 0, it will be set to 1 second.
func NewMonitor(ctx context.Context, every time.Duration) *Monitor {
	if every <= 0 {
		every = time.Second
	}
	ctr := &Monitor{
		C: make(chan int64),
	}
	go func(ctx context.Context) {
		ticker := time.NewTicker(every)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				select {
				case <-ctx.Done():
					return
				case ctr.C <- ctr.count.Swap(0):
				default:
				}
			case <-ctx.Done():
				return
			}
		}
	}(ctx)
	return ctr
}

// Rate increments the count and returns the current count.
func (c *Monitor) Rate() int64 {
	return c.count.Add(1)
}

// Value returns the current count.
func (c *Monitor) Value() int64 {
	return c.count.Load()
}
