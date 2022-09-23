package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

// Create a new method for the money type called printStr that returns the money value as a string with 2 decimal points.
func (m money) printStr() string {
	return fmt.Sprintf("%2f", m)
}

func main() {
	var salary money = 1.5 * 5.503
	fmt.Println(salary)  // result >>> 8.2545

	salary.print() // result >>> 8.25

	s:= salary.printStr()
	fmt.Println(s) // result >>> 8.254500

}