package main

import (
	"bufio"
	"os"
	"time"
)

// Version contains the binary version. This is added at build time.
var Version = "uncommitted"

// WantsColors represents whether the output should contain colors if the time
// between two subsequent lines "took too long".
// This is automatically toggled.
var WantsColors = false

// WantsBuffered represents whether the user would prefer output to be buffered,
// which makes it be processed a bit faster, or not.
// Running:
//   yes | head -n ... | tstdin
// Unbuffered processes ~1.1 Mb/s
// Buffered   processes ~1.4 Mb/s
var WantsBuffered = false

// FlushEvery represents, IFF WantsBuffered is set, after how many milliseconds
// STDOUT gets flushed.
var FlushEvery = 500

// ErrorAfter represents after how many ms to mark a line as having been output
// after "way too long". If colors are enabled, this is shown in red.
var ErrorAfter int = 60000

// WarnAfter represents after how many ms to mark a line as having been output
// after "a bit too long". If colors are enabled, this is whoen in yellow.
var WarnAfter int = 1000

func main() {
	dealWithArgs()
	if WantsBuffered {
		var stdOut = bufio.NewWriterSize(os.Stdout, 8192)
		ticker := time.NewTicker(time.Duration(FlushEvery) * time.Millisecond)
		done := make(chan bool)
		go func() {
			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					stdOut.Flush()
				}
			}
		}()
		timestamp(realClock{}, os.Stdin, stdOut, WantsColors)
		ticker.Stop()
		done <- true
		stdOut.Flush()
	} else {
		timestamp(realClock{}, os.Stdin, os.Stdout, WantsColors)
	}
}
