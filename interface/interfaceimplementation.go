package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {

	return r.width * r.height
}

func (c circle) area() float64 {

	return math.Pi * c.radius * c.radius
}

func (r rect) perim() float64 {

	return 2*r.width + 2*r.height

}

func (c circle) perim() float64 {

	return 2 * math.Pi * c.radius
}

func main() {

	r := rect{3, 4}
	c := circle{5}

	fmt.Println(r.area())
	fmt.Println(r.perim())
	fmt.Println(c.area())
	fmt.Println(c.perim())

}
