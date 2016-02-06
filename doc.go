// Package throttled offers wrappers for io.Reader and io.Writer that
// limit read/write speed by introducing delays in Read()/Write()
// calls.
//
// A speed limit is expressed with a Limiter, whose total available
// bandwidth can be shared between any number of readers and writers.
//
// A Limiter can also be used to rate-limit other actions, like
// sending email messages or making database queries.
//
// The bps program is an example: it copies stdin to stdout without
// exceeding a specified speed.
package throttled
