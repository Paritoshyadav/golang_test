package main

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {

	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, c := range cases {

		t.Run(testName(c.time), func(t *testing.T) {

			got := SecondHandInRadian(c.time)

			if got != c.angle {

				t.Errorf("got %v but want %v", got, c.angle)

			}

		})

	}

}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}
	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("Wanted %v Point, but got %v", c.point, got)
			}
		})
	}

}

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	want := Point{x: 150, y: 150 + 90}
	got := secondHand(tm)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

func simpleTime(hour, min, sec int) time.Time {
	return time.Date(1996, time.September, 14, hour, min, sec, 0, time.UTC)

}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func roughltEqual(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold

}

func roughlyEqualPoint(a, b Point) bool {
	return roughltEqual(a.x, b.x) &&
		roughltEqual(a.y, b.y)
}

// want := Point{x: 150, y: 150 - 90}
// got := SecondHand(cases.time)
// if got != want {
// 	t.Errorf("got %v but want %v", got, want)
