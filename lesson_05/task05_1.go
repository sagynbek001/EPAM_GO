package main

import (
	"fmt"

	interfaces "github.com/sagynbek001/EPAM_GO/interfaces"
	shapes "github.com/sagynbek001/EPAM_GO/shapes"
)

func DescribeShape(s interfaces.Shape) {
	fmt.Println(s)
	areaVal, _ := s.Area()
	perimeterVal, _ := s.Perimeter()
	fmt.Printf("Area: %.2f\n", areaVal)
	fmt.Printf("Perimeter: %.2f\n", perimeterVal)
}

func main() {
	// choose your own dimensions
	c := shapes.Circle{Radius: 10}
	r := shapes.Rectangle{
		Height: 3,
		Width:  6,
	}
	DescribeShape(c)
	DescribeShape(r)
}
