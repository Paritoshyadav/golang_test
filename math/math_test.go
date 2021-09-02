package main

import (
	"bytes"
	"encoding/xml"
	"math"
	"testing"
	"time"
)

type Svg struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

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

func TestSvgWriter(t *testing.T) {
	time := simpleTime(0, 0, 0)

	c := bytes.Buffer{}

	svgWriter(&c, time)
	svg := Svg{}

	xml.Unmarshal(c.Bytes(), &svg)
	want := Line{150, 150, 150, 60}
	x2 := "150.000"
	y2 := "60.000"

	for _, line := range svg.Line {
		if line == want {
			return
		}
	}

	t.Errorf("Expected to find the second hand with x2 of %+v and y2 of %+v, in the SVG output %v", x2, y2, svg.Line)

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
