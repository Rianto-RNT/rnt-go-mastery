package main

import "fmt"

type duration int

func main() {
	var hour duration = 3600
	minute := 60
	// different names are different types
	// we convert minute which is of type int to type duration
	fmt.Println(hour != duration(minute))
}
