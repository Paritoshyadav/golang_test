package main

import "time"

type Point struct {
	x float64
	y float64
}

func SecondHand(tm time.Time) Point {
	return Point{150, 60}
}

func main() {

}
