package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	"github.com/s4kh/learn-go-tests/util"
)

const (
	write = "write"
	sleep = "sleep"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buf := &bytes.Buffer{}
		spy := &SpySleeper{}

		Countdown(buf, spy)

		got := buf.String()
		want := "3\n2\n1\nGo!"

		util.AssertString(t, got, want)

		if spy.Calls != 3 {
			t.Errorf("not enough calls to sleeper, got %d want 3", spy.Calls)
		}
	})

	t.Run("sleep before every print after first print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}

		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted %v calls, got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
