/*
   Copyright 2016 Tom Clegg

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

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
// exceeding a specified speed.  See
// https://godoc.org/github.com/tomclegg/throttled/bps
//
//   go get github.com/tomclegg/throttled/bps
//   find ~ | bps -bytes 1 -per 3.3ms
//
// License
//
// Apache-2.0.
package throttled
