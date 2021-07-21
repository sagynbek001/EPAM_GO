package interfaces

type Shape interface {
	// add appropriate methods
	Area() (float64, error)
	Perimeter() (float64, error)
}
