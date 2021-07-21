package shapes

import (
	"errors"
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() (float64, error) {
	if c.Radius < 0 {
		return 0, errors.New("dimensions cannot be negative")
	}
	if c.Radius == 0 {
		return 0, errors.New("dimensions cannot be zero")
	}
	return math.Pi * math.Pow(c.Radius, 2), nil
}

func (c Circle) Perimeter() (float64, error) {
	if c.Radius < 0 {
		return 0, errors.New("dimensions cannot be negative")
	}
	if c.Radius == 0 {
		return 0, errors.New("dimensions cannot be zero")
	}
	return 2 * math.Pi * c.Radius, nil
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle: radius %.2f", c.Radius)
}
