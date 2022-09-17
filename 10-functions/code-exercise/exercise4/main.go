package main

import "fmt"

func sum(a ...int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func main() {
	s := sum(1, 2, 30)
	fmt.Println(s)
}
