package throttled

import (
	"gopkg.in/check.v1"
	"testing"
	"time"
)

func Test(t *testing.T) { check.TestingT(t) }

type Suite struct{}

var _ = check.Suite(&Suite{})

func (s *Suite) TestLimiter(c *check.C) {
	t := time.Now()
	l := NewLimiter(100, 10*time.Millisecond)
	l.Wait(10)
	c.Check(time.Since(t) < time.Millisecond, check.Equals, true)
	l.Wait(100)
	c.Check(time.Since(t) < time.Millisecond, check.Equals, true)
	l.Wait(10)
	c.Check(time.Since(t) > 10*time.Millisecond, check.Equals, true)
}

func ExampleLimiter_Wait() {
	// allow 100 units per millisecond
	l := NewLimiter(100, time.Millisecond)

	// these calls will return immediately:
	l.Wait(10)
	l.Wait(80)

	// this call will not return until the limiter is 1ms old:
	l.Wait(11)

	// this call will not return until the limiter is 20ms old:
	l.Wait(1000)

	time.Sleep(time.Second)

	// this call will block for ~1ms:
	l.Wait(100)
}
