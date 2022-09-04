package main

import "fmt"

func main() {

	const (
		//iota starts from zero
		jun = iota + 6
		jul //automatically incremented by one
		aug
	)

	fmt.Println(jun, jul, aug)
}
