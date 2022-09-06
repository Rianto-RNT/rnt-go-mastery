package main

import "fmt"

func main() {
	outer := 19
	_ = outer // blank identifier
	people := [5]string{"Rian", "Ryan", "Ian", "Rianto", "RNT"}
	friends := [2]string{"Saprol", "Udin"}

outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {

				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				break outer
			}
		}
	}
	fmt.Println("Next instruction after the break")
}
