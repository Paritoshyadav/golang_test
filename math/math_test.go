package main

import (
	"testing"
	"time"
)

func TestMath(t *testing.T) {

	tm := time.Date(1996, time.September, 14, 0, 0, 0, 0, time.UTC)
	want := Point{x: 150, y: 150 - 90}
	got := SecondHand(tm)
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}

}
