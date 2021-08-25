package main

import (
	"math"
	"time"
)

type Point struct {
	x float64
	y float64
}

//60 min = 360 , 1 min = 360/60 = 2*pi/60 = 2*pi/2*30=Pi/30
func SecondHandInRadian(tm time.Time) float64 {
	return (math.Pi / (30 / (float64(tm.Second()))))
}

func secondHandPoint(tm time.Time) Point {
	angle := SecondHandInRadian(tm)
	return Point{x: math.Sin(angle), y: math.Cos(angle)}
}

// 1) Scale it to the length of the hand
// 2) Flip it over the X axis because to account for the SVG having an origin in
// the top left hand corner
// 3) Translate it to the right position (so that it's coming from an origin of
// (150,150))
func secondHand(tm time.Time) Point {
	p := secondHandPoint(tm)
	p = Point{x: p.x * 90, y: p.y * 90}   //scale
	p = Point{x: p.x, y: -p.y}            //flip
	p = Point{x: p.x + 150, y: p.y + 150} //translate
	return p
}

func main() {

}
