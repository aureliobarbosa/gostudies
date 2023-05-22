package main

import "fmt"

type Triangle struct {
	height float64
	base   float64
}

func (t Triangle) getArea() float64 {
	return t.base * t.height / 2.0
}

type Square struct {
	sideLength float64
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type Shape interface {
	getArea() float64
}

func printArea(s Shape) {
	fmt.Println(s.getArea())
}

func main() {
	// var t Triangle

	t := Triangle{
		height: 5.0,
		base:   10.0,
	}
	s := Square{
		sideLength: 5.0,
	}

	fmt.Printf("Triangle Area: ")
	printArea(t)

	fmt.Printf("Square Area: ")
	printArea(s)
}
