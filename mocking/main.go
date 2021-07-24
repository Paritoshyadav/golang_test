package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const write = "write"
const sleep = "sleep"

type SpyTimer struct {
	sleep time.Duration
}

func (s *SpyTimer) Sleep(duration time.Duration) {
	s.sleep = duration
}

type ConfigTimeout struct {
	time  time.Duration
	sleep func(time.Duration)
}

func (c *ConfigTimeout) Sleep() {
	c.sleep(c.time)
}

type Sleeper interface {
	Sleep()
}

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
