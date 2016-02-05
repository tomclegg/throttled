package throttled

import (
	"io"
)

type reader struct {
	L *Limiter
	R io.Reader
}

func Reader(r io.Reader, lim *Limiter) io.ReadCloser {
	return &reader{
		L: lim,
		R: r,
	}
}

func (r *reader) Read(p []byte) (n int, err error) {
	n, err = r.R.Read(p)
	r.L.Wait(uint64(n))
	return
}

func (r *reader) Close() error {
	if c, ok := r.R.(io.Closer); ok {
		return c.Close()
	}
	return nil
}
