package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2. * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rectangle{
		width:  3,
		height: 4,
	}

	c := circle{
		radius: 5,
	}

	measure(r)
	measure(c)

}
