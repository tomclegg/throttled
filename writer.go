package throttled

import (
	"io"
)

type writer struct {
	L *Limiter
	W io.Writer
}

func Writer(w io.Writer, lim *Limiter) io.WriteCloser {
	return &writer{
		L: lim,
		W: w,
	}
}

func (w *writer) Write(p []byte) (n int, err error) {
	n, err = w.W.Write(p)
	w.L.Wait(uint64(n))
	return
}

func (w *writer) Close() error {
	if c, ok := w.W.(io.Closer); ok {
		return c.Close()
	}
	return nil
}
