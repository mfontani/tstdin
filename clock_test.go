package main

import (
	"testing"
)

func TestMonotonicClock(t *testing.T) {
	c := MonotonicClock(0, 0, 0, 5_000_000)
	now1 := c.Now()
	now2 := c.Now()
	if now1.Unix() == now2.Unix() && now1.UnixNano() == now2.UnixNano() {
		t.Fail()
		t.Logf("now1==now2 = %v & %v & %v.%v", now1, now2, now1.Unix(), now1.UnixNano())
		return
	}
	if now2.Unix() != now2.Unix() {
		t.Fail()
		t.Logf("now1.Unix()==now2.Unix() = %v & %v", now1, now2)
		return
	}
	if now2.UnixNano() != (now1.UnixNano() + 5_000_000) {
		t.Fail()
		t.Logf("now2.UnixNano()!=now1.UnixNano()+5_000_000 = now1:%v (%d.%d) & now2:%v (%d.%d) => %d!=%d",
			now1, now1.Unix(), now1.UnixNano(),
			now2, now2.Unix(), now2.UnixNano(),
			now2.UnixNano(), now1.UnixNano()+5_000_000,
		)
		return
	}
}
