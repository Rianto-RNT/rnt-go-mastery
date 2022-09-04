package main

import "fmt"

func main() {

	const (
		secPreDay = 60 * 60 * 24
		dayYear   = 365
	)

	fmt.Printf("There are %d seconds in a year.\n", secPreDay*dayYear)
}
