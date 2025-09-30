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
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(b, i)
		s.Sleep()
	}
	fmt.Fprintf(b, finalPhrase)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}
