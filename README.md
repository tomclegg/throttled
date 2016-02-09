# throttled
--
    import "github.com/tomclegg/throttled"

Package throttled offers wrappers for io.Reader and io.Writer that limit
read/write speed by introducing delays in Read()/Write() calls.

A speed limit is expressed with a Limiter, whose total available bandwidth can
be shared between any number of readers and writers.

A Limiter can also be used to rate-limit other actions, like sending email
messages or making database queries.

The bps program is an example: it copies stdin to stdout without exceeding a
specified speed. See https://godoc.org/github.com/tomclegg/throttled/bps

    go get github.com/tomclegg/throttled/bps
    find ~ | bps -bytes 1 -per 3.3ms


### License

Apache-2.0.

## Usage

#### func  Reader

```go
func Reader(r io.Reader, lim *Limiter) io.ReadCloser
```

#### func  Writer

```go
func Writer(w io.Writer, lim *Limiter) io.WriteCloser
```

#### type Limiter

```go
type Limiter struct {
}
```


#### func  NewLimiter

```go
func NewLimiter(flow uint64, interval time.Duration) *Limiter
```
NewLimiter returns a new Limiter that accepts the given flow in the given time
interval. It can be shared by multiple goroutines; they will share the available
flow.

#### func (*Limiter) Close

```go
func (lim *Limiter) Close()
```
Close frees all resources.

#### func (*Limiter) Wait

```go
func (lim *Limiter) Wait(flow uint64)
```
Wait until we're not over-limit -- i.e., until enough time has passed to admit
all previous Waits -- then count flow and return.
