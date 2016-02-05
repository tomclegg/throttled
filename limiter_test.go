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
