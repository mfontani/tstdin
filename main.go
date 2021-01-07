package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Version contains the binary version. This is added at build time.
var Version = "uncommitted"

// WantsColors represents whether the output should contain colors if the time
// between two subsequent lines "took too long".
// This is automatically toggled.
var WantsColors = false

const resetColor = "\033[0m"
const warningColor = "\033[33m"
const errorColor = "\033[31m"

// Neatly show a duration in hours, mins, seconds and microseconds.
func niceDuration(d time.Duration) string {
	return fmt.Sprintf("%02d:%02d:%02d.%06d",
		int(d.Hours()),
		int(d.Minutes())%60,
		int(d.Seconds())%60,
		d.Microseconds()%1000000,
	)
}

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
	if len(os.Args) == 2 {
		arg := os.Args[1]
		if arg == "-color" || arg == "--color" {
			WantsColors = true
		} else if arg == "-nocolor" || arg == "--nocolor" {
			WantsColors = false
		} else if arg == "-no-color" || arg == "--no-color" {
			WantsColors = false
		}
	}
}

func dealWithArgs() {
	// Support --version & --help only
	if len(os.Args) > 1 {
		wantsVersion := false
		wantsHelp := false
		for _, v := range os.Args[1:] {
			if v == "-version" || v == "--version" {
				wantsVersion = true
			}
			if v == "-help" || v == "--help" {
				wantsHelp = true
			}
		}
		if wantsVersion || wantsHelp {
			fmt.Printf("%s\n", Version)
			if wantsHelp {
				fmt.Println("")
			}
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
		}
		if wantsVersion || wantsHelp {
			os.Exit(0)
		}
	}
	if len(os.Args) != 1 {
		if len(os.Args) == 2 {
			arg := os.Args[1]
			// These args are okay:
			if arg == "-color" || arg == "--color" {
				return
			}
			if arg == "-nocolor" || arg == "--nocolor" {
				return
			}
			if arg == "-no-color" || arg == "--no-color" {
				return
			}
		}
		fmt.Fprintf(os.Stderr, "This command takes no arguments. See --help.\n")
		os.Exit(1)
	}
}

func timestampStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	startTime := time.Now()
	lastLine := time.Now()
	for scanner.Scan() {
		nowTime := time.Now()
		sinceStart := nowTime.Sub(startTime)
		sinceLastLine := nowTime.Sub(lastLine)
		color := ""
		reset := ""
		if WantsColors {
			if sinceLastLine.Minutes() > 1 {
				color = errorColor
				reset = resetColor
			} else if sinceLastLine.Seconds() > 1 {
				color = warningColor
				reset = resetColor
			}
		}
		fmt.Printf("%s %s %s%s%s %s\n",
			nowTime.Format("2006-01-02 15:04:05.000000"),
			niceDuration(sinceStart),
			color,
			niceDuration(sinceLastLine),
			reset,
			scanner.Text())
		lastLine = nowTime
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func main() {
	dealWithArgs()
	setWantsColors()
	timestampStdin()
}
