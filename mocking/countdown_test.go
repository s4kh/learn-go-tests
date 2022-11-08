package mocking

import (
	"bytes"
	"testing"

	"github.com/s4kh/learn-go-tests/util"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buf := &bytes.Buffer{}
	spy := &SpySleeper{}

	Countdown(buf, spy)

	got := buf.String()
	want := "3\n2\n1\nGo!"

	util.AssertString(t, got, want)

	if spy.Calls != 3 {
		t.Errorf("not enough calls to sleeper, got %d want 3", spy.Calls)
	}
}
