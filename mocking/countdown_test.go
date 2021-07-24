package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {

	t.Run("checking countdown output", func(t *testing.T) {
		count := &bytes.Buffer{}
		s := SpyCountSleeper{}
		Countdown(count, &s)

		got := count.String()

		want := "321go"

		assertCountdown(t, got, want)

	})
	t.Run("checking countdown sequence", func(t *testing.T) {

		s := SpyCountSleeper{}
		Countdown(&s, &s)

		got := s.calls

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		assertCountdown(t, got, want)

	})
	t.Run("checking countdown time response", func(t *testing.T) {
		sec := 5 * time.Second
		spy_timer := &SpyTimer{}
		configtime := ConfigTimeout{time: sec, sleep: spy_timer.Sleep}
		configtime.Sleep()

		if spy_timer.sleep != sec {
			t.Errorf("should have slept for %v but slept for %v", sec, spy_timer.sleep)
		}

	})

}

func assertCountdown(t testing.TB, got interface{}, want interface{}) {

	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

}
