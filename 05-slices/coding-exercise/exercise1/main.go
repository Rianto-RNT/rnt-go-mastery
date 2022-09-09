package main

import "fmt"

func main() {
	countries := []string{"Indonesia", "Sumbawa", "Lenangguar"}

	for i, v := range countries {
		fmt.Printf("Index: %d, Element: %q\n", i, v)
	}
}
