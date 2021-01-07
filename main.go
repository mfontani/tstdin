package main

import (
	"os"
)

// Version contains the binary version. This is added at build time.
var Version = "uncommitted"

// WantsColors represents whether the output should contain colors if the time
// between two subsequent lines "took too long".
// This is automatically toggled.
var WantsColors = false

func main() {
	dealWithArgs()
	stdOut := os.Stdout
	timestamp(realClock{}, os.Stdin, stdOut, WantsColors)
}
