package main

import "fmt"

type shape interface {
	area() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base float64
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{
		sideLength: 20,
	}

	t := triangle{
		height: 12,
		base: 8,
	}

	printArea(s)
	printArea(t)
}