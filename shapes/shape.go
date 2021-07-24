package shapes

import "math"

type Rectangle struct {
	length float64
	width  float64
}
type Circle struct {
	radius float64
}

//Reactangle Area Method

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {

	return 2 * (r.length + r.width)

}

//Circle Area Method
func (c Circle) Area() float64 {

	return math.Pi * math.Pow(c.radius, 2)

}
