package main

import (
	"fmt"
	"io"
	"iter"
	"os"
	"time"
)

const (
	finalPhrase   = "Go!"
	countdownFrom = 3
)

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type Sleeper interface {
	Sleep()
}

func Countdown(b io.Writer, s Sleeper) {
	for i := range countDownFrom(countdownFrom) {
		fmt.Fprintln(b, i)
		s.Sleep()
	}
	fmt.Fprintf(b, finalPhrase)
}

func countDownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}
