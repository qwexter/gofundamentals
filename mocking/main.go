package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalPhrase    = "Go!"
	countdownStart = 3
)

func main() {
	s := &realSleeper{}

	Countdown(os.Stdout, s)
}

type realSleeper struct{}

func (s *realSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type Sleeper interface {
	Sleep()
}

func Countdown(b io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(b, i)
		s.Sleep()
	}
	fmt.Fprintf(b, finalPhrase)
}
