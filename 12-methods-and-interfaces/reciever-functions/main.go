package main

import (
	"fmt"
	"time"
)

type names []string

func (n names) print() {
	for i, names := range n {
		fmt.Println(i, names)
	}
}

func main() {
	const day = 24 * time.Hour
	fmt.Printf("%T\n", day)

	seconds := day.Seconds()

	fmt.Printf("%T\n", seconds)
	fmt.Printf("Seconds in a day: %v\n", seconds)

	friends := names{"Rian", "Morrison"}
	friends.print()

	names.print(friends)

	var n int64 = 7451155522
	fmt.Println(n)
	fmt.Println(time.Duration(n))
}