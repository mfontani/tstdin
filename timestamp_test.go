package main

import (
	"bytes"
	"os"
	"testing"
)

var fakeEpoch int64 = 909090909

// perl -lE'say scalar gmtime 909090909'
// Thu Oct 22 21:15:09 1998

func TestTimestampSingleLine(t *testing.T) {
	if os.Getenv("TZ") != "UTC" {
		t.Fail()
		t.Logf("Run tests with TZ=UTC!")
		return
	}
	c := StuckClock(fakeEpoch, 0)
	in := bytes.NewBufferString("foo")
	var outB bytes.Buffer
	timestamp(c, in, &outB, false)
	wantS := "1998-10-22 21:15:09.000000 00:00:00.000000 00:00:00.000000 foo\n"
	if outB.String() != wantS {
		t.Fail()
		t.Logf("GOT  bytes: %v", outB)
		t.Logf("WANT string:\n%v", wantS)
		t.Logf("GOT  string:\n%v", outB.String())
	}
}

func TestTimestampTwoLines(t *testing.T) {
	if os.Getenv("TZ") != "UTC" {
		t.Fail()
		t.Logf("Run tests with TZ=UTC!")
		return
	}
	c := StuckClock(fakeEpoch, 0)
	in := bytes.NewBufferString("foo\nbar")
	var outB bytes.Buffer
	timestamp(c, in, &outB, false)
	wantS := "1998-10-22 21:15:09.000000 00:00:00.000000 00:00:00.000000 foo\n1998-10-22 21:15:09.000000 00:00:00.000000 00:00:00.000000 bar\n"
	if outB.String() != wantS {
		t.Fail()
		t.Logf("GOT  bytes: %v", outB)
		t.Logf("WANT string:\n%v", wantS)
		t.Logf("GOT  string:\n%v", outB.String())
	}
}

func TestTimestampTwoLinesColored(t *testing.T) {
	if os.Getenv("TZ") != "UTC" {
		t.Fail()
		t.Logf("Run tests with TZ=UTC!")
		return
	}
	c := StuckClock(fakeEpoch, 0)
	in := bytes.NewBufferString("foo\nbar")
	var outB bytes.Buffer
	timestamp(c, in, &outB, true)
	wantS := "1998-10-22 21:15:09.000000 00:00:00.000000 00:00:00.000000 foo\n1998-10-22 21:15:09.000000 00:00:00.000000 00:00:00.000000 bar\n"
	if outB.String() != wantS {
		t.Fail()
		t.Logf("GOT  bytes: %v", outB)
		t.Logf("WANT string:\n%v", wantS)
		t.Logf("GOT  string:\n%v", outB.String())
	}
}

func TestTimestampTwoLinesHalfSecond(t *testing.T) {
	if os.Getenv("TZ") != "UTC" {
		t.Fail()
		t.Logf("Run tests with TZ=UTC!")
		return
	}
	c := MonotonicClock(fakeEpoch, 0, 0, 500_000_000)
	in := bytes.NewBufferString("foo\nbar")
	var outB bytes.Buffer
	timestamp(c, in, &outB, false)
	wantS := "1998-10-22 21:15:09.500000 00:00:00.500000 00:00:00.500000 foo\n1998-10-22 21:15:10.000000 00:00:01.000000 00:00:00.500000 bar\n"
	if outB.String() != wantS {
		t.Fail()
		t.Logf("GOT  bytes: %v", outB)
		t.Logf("WANT string:\n%v", wantS)
		t.Logf("GOT  string:\n%v", outB.String())
	}
}

func TestTimestampTwoLinesHalfSecondColored(t *testing.T) {
	if os.Getenv("TZ") != "UTC" {
		t.Fail()
		t.Logf("Run tests with TZ=UTC!")
		return
	}
	c := MonotonicClock(fakeEpoch, 0, 0, 500_000_000)
	in := bytes.NewBufferString("foo\nbar")
	var outB bytes.Buffer
	timestamp(c, in, &outB, true)
	wantS := "1998-10-22 21:15:09.500000 00:00:00.500000 00:00:00.500000 foo\n1998-10-22 21:15:10.000000 00:00:01.000000 00:00:00.500000 bar\n"
	if outB.String() != wantS {
		t.Fail()
		t.Logf("GOT  bytes: %v", outB)
		t.Logf("WANT string:\n%v", wantS)
		t.Logf("GOT  string:\n%v", outB.String())
	}
}

func TestTimestampTwoLinesOneHalfSecond(t *testing.T) {
	if os.Getenv("TZ") != "UTC" {
		t.Fail()
		t.Logf("Run tests with TZ=UTC!")
		return
	}
	c := MonotonicClock(fakeEpoch, 0, 1, 500_000_000)
	in := bytes.NewBufferString("foo\nbar")
	var outB bytes.Buffer
	timestamp(c, in, &outB, false)
	wantS := "1998-10-22 21:15:10.500000 00:00:01.500000 00:00:01.500000 foo\n1998-10-22 21:15:12.000000 00:00:03.000000 00:00:01.500000 bar\n"
	if outB.String() != wantS {
		t.Fail()
		t.Logf("GOT  bytes: %v", outB)
		t.Logf("WANT string:\n%v", wantS)
		t.Logf("GOT  string:\n%v", outB.String())
	}
}

func TestTimestampTwoLinesOneHalfSecondColored(t *testing.T) {
	if os.Getenv("TZ") != "UTC" {
		t.Fail()
		t.Logf("Run tests with TZ=UTC!")
		return
	}
	c := MonotonicClock(fakeEpoch, 0, 1, 500_000_000)
	in := bytes.NewBufferString("foo\nbar")
	var outB bytes.Buffer
	timestamp(c, in, &outB, true)
	wantS := "1998-10-22 21:15:10.500000 00:00:01.500000 \033[33m00:00:01.500000\033[0m foo\n1998-10-22 21:15:12.000000 00:00:03.000000 \033[33m00:00:01.500000\033[0m bar\n"
	if outB.String() != wantS {
		t.Fail()
		t.Logf("WANT bytes: %v", []byte(wantS))
		t.Logf("GOT  bytes: %v", outB)
		t.Logf("WANT string:\n%v", wantS)
		t.Logf("GOT  string:\n%v", outB.String())
	}
}

func TestTimestampTwoLinesNinetySeconds(t *testing.T) {
	if os.Getenv("TZ") != "UTC" {
		t.Fail()
		t.Logf("Run tests with TZ=UTC!")
		return
	}
	c := MonotonicClock(fakeEpoch, 0, 90, 0)
	in := bytes.NewBufferString("foo\nbar")
	var outB bytes.Buffer
	timestamp(c, in, &outB, false)
	wantS := "1998-10-22 21:16:39.000000 00:01:30.000000 00:01:30.000000 foo\n1998-10-22 21:18:09.000000 00:03:00.000000 00:01:30.000000 bar\n"
	if outB.String() != wantS {
		t.Fail()
		t.Logf("GOT  bytes: %v", outB)
		t.Logf("WANT string:\n%v", wantS)
		t.Logf("GOT  string:\n%v", outB.String())
	}
}

func TestTimestampTwoLinesNinetySecondsColored(t *testing.T) {
	if os.Getenv("TZ") != "UTC" {
		t.Fail()
		t.Logf("Run tests with TZ=UTC!")
		return
	}
	c := MonotonicClock(fakeEpoch, 0, 90, 0)
	in := bytes.NewBufferString("foo\nbar")
	var outB bytes.Buffer
	timestamp(c, in, &outB, true)
	wantS := "1998-10-22 21:16:39.000000 00:01:30.000000 \033[31m00:01:30.000000\033[0m foo\n1998-10-22 21:18:09.000000 00:03:00.000000 \033[31m00:01:30.000000\033[0m bar\n"
	if outB.String() != wantS {
		t.Fail()
		t.Logf("GOT  bytes: %v", outB)
		t.Logf("WANT string:\n%v", wantS)
		t.Logf("GOT  string:\n%v", outB.String())
	}
}
