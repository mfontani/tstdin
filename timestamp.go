package main

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

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

func timestamp(clock Time, r io.Reader, w io.Writer, wantsColors bool) {
	scanner := bufio.NewScanner(r)
	startTime := clock.Now()
	lastLine := startTime
	for scanner.Scan() {
		nowTime := clock.Now()
		sinceStart := nowTime.Sub(startTime)
		sinceLastLine := nowTime.Sub(lastLine)
		color := ""
		reset := ""
		if wantsColors {
			if sinceLastLine.Milliseconds() > int64(ErrorAfter) {
				color = errorColor
				reset = resetColor
			} else if sinceLastLine.Milliseconds() > int64(WarnAfter) {
				color = warningColor
				reset = resetColor
			}
		}
		fmt.Fprintf(w, "%s %s %s%s%s %s\n",
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
