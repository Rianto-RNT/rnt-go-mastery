package main

import "fmt"

func sum(a ...int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

func main() {
	s := sum(1, 2, 30)
	fmt.Println(s)
}
