package throttled

import (
	"time"
)

type Limiter struct {
	flow     uint64
	interval time.Duration
	c        chan uint64
	done     chan struct{}
}

// NewLimiter returns a new Limiter that accepts the given flow in the
// given time interval. It can be shared by multiple goroutines; they
// will share the available flow.
func NewLimiter(flow uint64, interval time.Duration) *Limiter {
	lim := &Limiter{
		flow:     flow,
		interval: interval,
		c:        make(chan uint64),
		done:     make(chan struct{}),
	}
	go lim.run()
	return lim
}

// Wait until we're not over-limit -- i.e., until enough time has
// passed to admit all previous Waits -- then count flow and return.
func (lim *Limiter) Wait(flow uint64) {
	lim.c <- flow
}

// Close frees all resources.
func (lim *Limiter) Close() {
	close(lim.c)
}

func (lim *Limiter) run() {
	t := time.NewTicker(lim.interval)
	var accum uint64
	for f := range lim.c {
		for accum += f; accum > lim.flow; accum -= lim.flow {
			<-t.C
		}
	}
	t.Stop()
}
