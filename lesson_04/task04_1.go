package main

import (
	"fmt"
	"math"
)

type Shape interface {
	// add appropriate methods
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle: radius %.2f", c.radius)
}

type Rectangle struct {
	height float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle with height %.2f and width %.2f", r.height, r.width)
}

func DescribeShape(s Shape) {
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	// choose your own dimensions
	c := Circle{radius: 10}
	r := Rectangle{
		height: 3,
		width:  6,
	}
	DescribeShape(c)
	DescribeShape(r)
}
