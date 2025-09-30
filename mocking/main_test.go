package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("when check the order of operations", func(t *testing.T) {
		ops := &SpyCountdownOperations{}

		Countdown(ops, ops)

		want := []string{
			writeOperation,
			sleepOperation,
			writeOperation,
			sleepOperation,
			writeOperation,
			sleepOperation,
			writeOperation,
		}

		if !reflect.DeepEqual(want, ops.Calls) {
			t.Errorf("want %v, got %v", want, ops.Calls)
		}
	})
	t.Run("when check the output string", func(t *testing.T) {
		ops := &SpyCountdownOperations{}
		b := &bytes.Buffer{}

		Countdown(b, ops)

		want := "3\n2\n1\nGo!"

		got := b.String()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleep := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := &ConfigurableSleeper{sleep, spyTime.SetDurationSlept}
	sleeper.Sleep()

	if spyTime.durationSlept != sleep {
		t.Errorf("should have slept for %v but slept for %v", sleep, spyTime.durationSlept)
	}
}

const (
	writeOperation = "write"
	sleepOperation = "sleep"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, writeOperation)
	return 0, nil
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleepOperation)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}
