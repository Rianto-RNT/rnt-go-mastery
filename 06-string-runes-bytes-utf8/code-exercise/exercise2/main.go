package main

import "fmt"

func main() {
	r := 'ă'                     // declaring a rune
	fmt.Printf("r type:%T\n", r) // rune is alias to int32

	s1, s2 := "ma", "m" // declaring 2 strings

	// concatenating strings
	s := s1 + s2 + string(r)   // converting rune to string (expliction conversion is required)
	fmt.Printf("s is %s\n", s) // => s is mamă

}
