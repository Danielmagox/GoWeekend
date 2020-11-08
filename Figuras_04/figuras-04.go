package main

import (
	"fmt"
	"math"
)

//Circle Structure
type Circle struct {
	perimeter float64
}

//Square Structure
type Square struct {
	side float64
}

//Triangle Structure
type Triangle struct {
	height float64
	base   float64
	kind   string
}

func (circ Circle) area() float64 {
	return circ.perimeter / 2 * math.Pi
}

func (square Square) area() float64 {
	return math.Pow(square.side, 2)
}

func (triangle Triangle) area() float64 {
	return triangle.base * triangle.height
}

func main() {
	circ := Circle{
		perimeter: 5.1212345234,
	}
	square := Square{
		side: 9.129391239123,
	}
	triangle := Triangle{
		height: 10.823482348,
		base:   29.12309123,
	}

	fmt.Printf("circle Area %f ", circ.area())
	fmt.Printf("square Area %f ", square.area())
	fmt.Printf("triangle Area %f ", triangle.area())
}
