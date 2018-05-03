package main

import (
	"math"
	"fmt"
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

type triangle struct {
	length  float64
	breadth float64
	height  float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (t triangle) area() float64 {
	return 0.5 * t.height * t.breadth
}
func (t triangle) perim() float64 {
	return 0.5 * t.height * t.breadth
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 2, height: 4}
	measure(r)

	c := circle{radius: 2}
	measure(c)

	t := triangle{1, 2, 3}
	measure(t)
}
