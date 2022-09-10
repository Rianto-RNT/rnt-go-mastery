package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Golang"
	fmt.Println(len(s1))

	name := "Cătălin"
	fmt.Println(len(name))

	fmt.Println(len("ريانتو")) // "ريانتو" = "rianto" via google translate

	n := utf8.RuneCountInString(name)
	m := utf8.RuneCountInString("ريانتو")
	fmt.Println("n:", n, "m:", m)

}
