package poc

import (
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.Radius
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// type Border interface {
// 	Perimeter() float64
// }

// func printBorder(b Border) {
// 	fmt.Println(b.Perimeter())
// }
