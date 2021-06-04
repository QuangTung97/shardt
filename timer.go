package shardt

import "time"

const hundredYears = 100 * 365 * 24 * time.Hour

type timer interface {
	reset(d time.Duration)
	resetAfterChan(d time.Duration)
	getChan() <-chan time.Time
}

type simpleTimer struct {
	timer *time.Timer
}

func newSimpleTimer() simpleTimer {
	return simpleTimer{
		timer: time.NewTimer(hundredYears),
	}
}

func (t simpleTimer) reset(d time.Duration) {
	if !t.timer.Stop() {
		<-t.timer.C
	}
	t.timer.Reset(d)
}

func (t simpleTimer) resetAfterChan(d time.Duration) {
	t.timer.Reset(d)
}

func (t simpleTimer) getChan() <-chan time.Time {
	return t.timer.C
}
