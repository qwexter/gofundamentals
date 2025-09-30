package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	s := &SpySleeper{}

	Countdown(buffer, s)

	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if s.Calls != 3 {
		t.Errorf("got %q, want %q", s.Calls, 3)
	}
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
