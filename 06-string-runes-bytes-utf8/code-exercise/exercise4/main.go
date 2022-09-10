package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	/** ERROR CODE
	  s1 := "Go is cool!"
	  x := s1[0]
	  fmt.Println(x)

	  s1[0] = 'x'

	  // printing the number of runes in the string
	  fmt.Println(len(s1))
	*/

	/*VALID CODE*/
	s1 := "Whatever i'm late anyway"

	x := s1[0]
	fmt.Println(x)

	// ERROR -> cannot assign to s1[0]
	// s1[0] = 'x'

	// printing the number of runes in the string
	// fmt.Println(len(s1))
	fmt.Println(utf8.RuneCountInString(s1))

}
