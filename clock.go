package main

import "time"

// Time is a custom interface to ensure we can test timestamp()
type Time interface {
	Now() time.Time
	Sub(time.Time) time.Duration
}

type realClock struct{}

func (realClock) Now() time.Time                  { return time.Now() }
func (c realClock) Sub(t time.Time) time.Duration { return c.Sub(t) }

type stuckClock struct {
	sec  int64
	nsec int64
}

func (c stuckClock) Now() time.Time {
	return time.Unix(c.sec, c.nsec)
}

func (c stuckClock) Sub(t time.Time) time.Duration {
	when := time.Unix(c.sec, c.nsec)
	return when.Sub(time.Unix(c.sec, c.nsec))
}

// StuckClock creates a new "stuck" clock, which starts at the given sec, nsec
// and always returns itself for any Sub() call
func StuckClock(sec, nsec int64) Time {
	return stuckClock{sec: sec, nsec: nsec}
}

type monotonicClock struct {
	sec          int64
	nsec         int64
	secIncrease  int64
	nsecIncrease int64
}

func (c *monotonicClock) Now() time.Time {
	then := time.Unix(c.sec, c.nsec)
	c.sec += c.secIncrease
	c.nsec += c.nsecIncrease
	for c.nsec > 1_000_000_000 {
		c.sec++
		c.nsec -= 1_000_000_000
	}
	return then
}

func (c monotonicClock) Sub(t time.Time) time.Duration {
	when := time.Unix(c.sec, c.nsec)
	return when.Sub(t)
}

// MonotonicClock creates a new "stuck" clock, which starts at the given sec,
// nsec and whenever Now() is called, it returns the last time incremented by
// the given delta seconds and nsec.
func MonotonicClock(sec, nsec, secIncrease, nsecIncrease int64) Time {
	mc := monotonicClock{sec: sec, nsec: nsec, secIncrease: secIncrease, nsecIncrease: nsecIncrease}
	return &mc
}
