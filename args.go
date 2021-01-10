package main

import (
	"fmt"
	"os"
	"strconv"
)

// automatically turn colors on/off depending on the circumstances in which
// we're being called.
func setWantsColors() {
	// If STDOUT is a terminal, we can have colors.
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) != 0 {
		WantsColors = true
	}
	// If the user's specifically requesting no colors from any application, we
	// should honor their choice.
	if os.Getenv("NO_COLOR") != "" {
		WantsColors = false
	}
	// If the user's specifically requesting no colors from THIS application, we
	// should honor their choice, too.
	if os.Getenv("TSTDIN_NO_COLOR") != "" {
		WantsColors = false
	}
	// A dumb terminal should never show colors.
	if os.Getenv("TERM") == "dumb" {
		WantsColors = false
	}
	// The user might want to override either behaviour
	overrodeColor := false
	overrodeNoColor := false
	for _, v := range os.Args[1:] {
		if v == "-color" || v == "--color" {
			WantsColors = true
			overrodeColor = true
		} else if v == "-nocolor" || v == "--nocolor" {
			WantsColors = false
			overrodeNoColor = true
		} else if v == "-no-color" || v == "--no-color" {
			WantsColors = false
			overrodeNoColor = true
		}
	}
	if overrodeColor && overrodeNoColor {
		fmt.Fprintf(os.Stderr, "Can't override both -color and -nocolor. See --help.\n")
		os.Exit(1)
	}
}

func dealWithArgs() {
	// Support --version & --help only
	if len(os.Args) > 1 {
		wantsVersion := false
		wantsHelp := false
		spuriousArgs := false
		skipArg := false
		for i, v := range os.Args[1:] {
			if skipArg {
				skipArg = false
				continue
			}
			if v == "-version" || v == "--version" {
				wantsVersion = true
			} else if v == "-help" || v == "--help" {
				wantsHelp = true
			} else if v == "-nocolor" || v == "--nocolor" || v == "-no-color" || v == "--no-color" {
			} else if v == "-color" || v == "--color" {
			} else if v == "-buffered" || v == "--buffered" {
				WantsBuffered = true
			} else if v == "-flushevery" || v == "--flushevery" || v == "-flush-every" || v == "--flush-every" {
				WantsBuffered = true
				if len(os.Args) > i {
					var err error
					FlushEvery, err = strconv.Atoi(os.Args[i+2])
					if err != nil {
						fmt.Fprintf(os.Stderr, "Bad argument %s=%s. Need a numeric value representing ms. See --help.\n", v, os.Args[i+2])
						os.Exit(1)
					}
					skipArg = true
				} else {
					fmt.Fprintf(os.Stderr, "Bad arguments. %s wants a value in ms. No argument given. See --help.\n", v)
					os.Exit(1)
				}
			} else {
				spuriousArgs = true
			}
		}
		if wantsVersion {
			fmt.Printf("%s\n", Version)
			os.Exit(0)
		}
		if wantsHelp {
			fmt.Println("Usage: tstdin [OPTIONS]")
			fmt.Println("")
			fmt.Println("Prints out lines received via STDIN, prepending:")
			fmt.Println("- (1) the current date/time including microseconds")
			fmt.Println("- (2) hours, minutes, seconds and microseconds since the start")
			fmt.Println("- (3) hours, minutes, seconds and microseconds since the last line was received")
			fmt.Println("- (4) the line that was received")
			fmt.Println("... all separated by spaces.")
			fmt.Println("")
			fmt.Println("If the time between the previous line and the current one took too long, the")
			fmt.Println("(3) token will be output with color: red if it took more than a minute, or")
			fmt.Println("yellow if it took more than a second.")
			fmt.Println("Colors are automatically disabled if STDOUT is not a terminal, if TERM=dumb")
			fmt.Println("or if either NO_COLOR or TSTDIN_NO_COLOR are set, but see -color and -no-color.")
			fmt.Println("")
			fmt.Println("Example:")
			fmt.Println("  $ (echo foo ; sleep 1; echo bar ; sleep 1 ; echo baz) | tstdin")
			fmt.Println("  2031-04-07 14:20:03.952221 00:00:00.000003 00:00:00.000003 foo")
			fmt.Println("  2031-04-07 14:20:04.952120 00:00:00.999902 00:00:00.999898 bar")
			fmt.Println("  2031-04-07 14:20:05.953362 00:00:02.001144 00:00:01.001241 baz")
			fmt.Println("")
			fmt.Println("Options:")
			fmt.Println("  -help     Shows this help page")
			fmt.Println("  -version  Shows the program's version")
			fmt.Println("  -color    Force color on regardless of environment")
			fmt.Println("  -no-color Force color off regardless of environment")
			fmt.Printf("  -buffered Buffers stdout for output (might be faster), flushing every %dms\n", FlushEvery)
			fmt.Printf("  -flush-every MS (default: %d)\n", FlushEvery)
			fmt.Println("            Choose how often stdout is flushed. Implies -buffered.")
			fmt.Printf("\nThis is tstdin %s\n", Version)
			os.Exit(0)
		}
		if spuriousArgs {
			fmt.Fprintf(os.Stderr, "Bad arguments. See --help.\n")
			os.Exit(1)
		}
	}
	setWantsColors()
}
