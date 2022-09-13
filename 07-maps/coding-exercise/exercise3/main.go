package main

import "fmt"

func main() {
	// 1.
	m := map[int]bool{1: true, 2: false, 3: false}

	// 2.
	delete(m, 2)

	// 3.
	for k, v := range m {
		fmt.Printf("k: %d, v: %t\n", k, v)
	}
}
