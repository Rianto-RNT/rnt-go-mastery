package main

import "fmt"

type duration int

func main() {
	var hour duration
	fmt.Printf("hour's type: %T, hour's value: %v\n", hour, hour)
	hour = 3600
	fmt.Printf("hour's value %v\n", hour)
}
