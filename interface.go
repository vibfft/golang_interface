package main

import (
	"fmt"
	"math"
)

// Shape interface has many methods
type Shape interface {
	shapeName() string
	area() float64
	perimeter() float64
}

// Circle has center point and radius
type Circle struct {
	x float64
	y float64
	r float64
}

// Rectangle has four points
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r Rectangle) shapeName() string {
	return "Rectangle"
}

func (c Circle) shapeName() string {
	return "Circle"
}

func (r Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2*l + 2*w
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func measure(s Shape) {
	fmt.Println(s.shapeName())
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {

	c := Circle{0, 0, 5}
	measure(c)

	r := Rectangle{0, 0, 10, 10}
	measure(r)

	fmt.Println("Total Area for ", c.shapeName(), " and ", r.shapeName())
	fmt.Println(totalArea(&c, &r))

}
