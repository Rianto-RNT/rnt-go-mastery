package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) parameter() float64{
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64{
	return r.height * r.width
}

func (r rectangle) parameter() float64{
	return 2 * (r.height * r.width)
}

func printCircle(c circle) {
	fmt.Println("Shape:", c)
	fmt.Println("Area:", c.area())
	fmt.Println("Parameter:", c.parameter())
}

func printRectangle(r rectangle){
	fmt.Println("Shape:", r)
	fmt.Println("Area:", r.area())
	fmt.Println("Parameter:", r.parameter())
}

func main() {
	c1 := circle{radius: 5.}
	r1 := rectangle{width: 3., height: 2.1}

	printCircle(c1)
	printRectangle(r1)
}
