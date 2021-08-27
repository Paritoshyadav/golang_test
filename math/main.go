package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"time"
)

const secondHandLength = 90
const clockCentreX = 150
const clockCentreY = 150

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`

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
	p = Point{x: p.x * secondHandLength, y: p.y * secondHandLength} //scale
	p = Point{x: p.x, y: -p.y}                                      //flip
	p = Point{x: p.x + clockCentreX, y: p.y + clockCentreY}         //translate
	return p
}
func secondHandTag(p Point) string {
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.x, p.y)

}

func svgWriter(w io.Writer, t time.Time) {
	sh := secondHand(t)
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	io.WriteString(w, secondHandTag(sh))
	io.WriteString(w, svgEnd)

}

func main() {
	time := time.Now()
	svgWriter(os.Stdout, time)

}
