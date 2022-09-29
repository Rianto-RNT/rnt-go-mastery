package main

import (
	"fmt"
	"rntMyGoPackages/numbers"
)

func main() {
	var x uint = 22
	fmt.Printf("%d is even: %t\n", x, numbers.Even(x))
}
