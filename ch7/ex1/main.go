package main

import (
	"fmt"
	"math"
)

// a better way to write a program is to use a struct rather than using only Go's built-in data types.
// a struct is a type that contains names 'fields'
// Rectangle and Circle have area methods that return float64s, so both types implement the Shape interface
// we can use interface types as arguments to functions

type Shape interface {
	// An interface is like struct that defines a 'method set' instead of a 'field'
	// A 'method set' is a list of methods that a type must have in order to implement the interface
	area() float64
	perimeter() float64
}

type Circle struct { // a struct is a type that contains names 'fields'. Fields are like a set off grouped variables.
	// each field has a name and a type and is stored adjacent to the other ffields in the struct.
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func main() {
	c := Circle{1, 1, 5} // create an instance of new Circle type
	r := Rectangle{1, 1, 10, 10}

	fmt.Println(c.area()) // . calls the Circle area function
	fmt.Println(r.area()) // . calls the Rectange area function

	fmt.Println(c.perimeter()) // . calls the Circle perimeter function
	fmt.Println(r.perimeter()) // . calls the Rectangle perimeter function

}

func (c *Circle) area() float64 { //special type of function known as a method. (c *Circle) is a receiver.
	// Allows us to call the function using the . operator:
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