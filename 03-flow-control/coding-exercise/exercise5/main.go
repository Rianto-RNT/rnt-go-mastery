package main

import "fmt"

func main() {
	birthYear, currentYear := 1993, 2022

	for i := birthYear; i <= currentYear; {
		fmt.Printf("%d\n", i)
		i++
	}
	fmt.Println()
}
