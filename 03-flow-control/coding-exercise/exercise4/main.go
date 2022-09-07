package main

import "fmt"

func main() {

	for i := 1; i <= 500; i++ {
		if i%7 == 0 && i%5 == 0 { // if i is divisible both by 7 and 5
			fmt.Printf("%d\n", i)
		}
	}
	fmt.Println("")
}
