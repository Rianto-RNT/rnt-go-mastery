package main

import "fmt"

func main() {

	for i := 1; i <= 50; i++ {
		if i%7 == 0 { // if i is divisible by 7
			fmt.Printf("%d\n", i)
		}
	}
	fmt.Println("")
}
