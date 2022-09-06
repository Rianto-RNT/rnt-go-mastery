package main

import "fmt"

func main() {
	// 1)
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)

	// }

	// 2)
	for i := 0; i < 10; {
		fmt.Println(i)
		i++
	}
}
