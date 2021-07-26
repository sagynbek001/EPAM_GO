package shapes

import (
	"errors"
	"fmt"
)

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() (float64, error) {
	if r.Height < 0 || r.Width < 0 {
		return 0, errors.New("dimensions cannot be negative")
	}
	if r.Height == 0 || r.Width == 0 {
		return 0, errors.New("dimensions cannot be zero")
	}
	return r.Height * r.Width, nil
}

func (r Rectangle) Perimeter() (float64, error) {
	if r.Height < 0 || r.Width < 0 {
		return 0, errors.New("dimensions cannot be negative")
	}
	if r.Height == 0 || r.Width == 0 {
		return 0, errors.New("dimensions cannot be zero")
	}
	return 2 * (r.Height + r.Width), nil
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle with height %.2f and width %.2f", r.Height, r.Width)
}
