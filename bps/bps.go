// bps copies stdin to stderr, without exceeding a specified rate.
//
// Emulate a 2400 baud modem:
//
//   $ find ~ | bps -bytes 1 -per 3.3ms
//
// Another pointless exercise:
//
//   $ dd if=/dev/zero bs=1000 count=12345 | bps -bytes 12345 -per 10ms >/dev/null
//   12345+0 records in
//   12345+0 records out
//   12345000 bytes (12 MB) copied, 10.4019 s, 1.2 MB/s
//
package main

import (
	"flag"
	"io"
	"log"
	"os"
	"time"

	"github.com/tomclegg/throttled"
)

func main() {
	flow := flag.Uint64("bytes", 16000, "max bytes to copy per time interval")
	dur := flag.Duration("per", time.Second, "time interval")
	flag.Parse()

	bufsize := 1<<22
	if *flow < uint64(bufsize) {
		bufsize = int(*flow)
	}

	_, err := io.CopyBuffer(
		os.Stdout,
		throttled.Reader(os.Stdin, throttled.NewLimiter(*flow, *dur)),
		make([]byte, bufsize))
	if err != nil {
		log.Fatal(err)
	}
}
