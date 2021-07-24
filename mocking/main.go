package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const write = "write"
const sleep = "sleep"

// for third test
type SpyTimer struct {
	sleep time.Duration
}

func (s *SpyTimer) Sleep(duration time.Duration) {
	s.sleep = duration
}

//time configuration structure which take time and sleep function and call that sleep function using Sleep method
type ConfigTimeout struct {
	time  time.Duration
	sleep func(time.Duration)
}

func (c *ConfigTimeout) Sleep() {
	c.sleep(c.time)
}

// Sleeper interface for testing
type Sleeper interface {
	Sleep()
}

// For second test to make sure that sleep sequence is correct
type SpyCountSleeper struct {
	calls []string
}

func (s *SpyCountSleeper) Sleep() {

	s.calls = append(s.calls, sleep)

}

func (s *SpyCountSleeper) Write(p []byte) (n int, err error) {
	s.calls = append(s.calls, write)
	return

}

//main countdown function
func Countdown(out io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		s.Sleep()
		fmt.Fprintf(out, "%d", i)
	}
	s.Sleep()
	fmt.Fprintf(out, "go")
}

func main() {
	s := ConfigTimeout{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, &s)

}
