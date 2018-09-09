package main

import (
	"fmt"
	"math"
)

// a better way to write a program is to use a struct
// a struct is a type that contains names 'fields'
// Rectangle and Circle have area methods that return float64

type Shape interface {
	// An interface is like struct that defines a 'method set' instead of a 'field'
	// A 'method set' is a list of methods that a type must have in order to implement the interface
	area() float64
	perimeter() float64
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func main() {
	c := Circle{1, 1, 5}
	r := Rectangle{1, 1, 10, 10}

	fmt.Println(c.area())
	fmt.Println(r.area())

	fmt.Println(c.perimeter())
	fmt.Println(r.perimeter())

}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)

}
