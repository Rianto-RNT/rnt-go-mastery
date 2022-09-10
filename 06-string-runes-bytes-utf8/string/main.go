package main

import "fmt"

func main() {
	// Strings are defined between double quotes "..."
	// Strings in Go are UTF-8 encoded by default
	// A string is in fact a slice of bytes in Go

	// declaring a string
	s1 := "Hi there, RNT!"

	// printing a string
	fmt.Printf("%s\n", s1) // Hi there, RNT!
	fmt.Printf("%q\n", s1) // "Hi there, RNT!"

	// using double-quotes inside double quotes
	fmt.Println("Good Morning: \"RNT\"") // Good Morning: "RNT"

	// double quotes inside backticks (backquote)
	fmt.Println(`Good Afternoon: "Rianto"`) // Good Afternoon: "Rianto"

	// a string literal inclosed in backticks is called a raw string and it is interpreted literally.
	// backslashes or \n  have no special meaning
	s2 := `Whatever im late anyway`
	fmt.Println(s2) // Whatever im late anyway

	// declaring a multiline string
	fmt.Println("Price: 150\nBrand: G-Shock")
	//the same with:
	fmt.Println(`
Price: 150
Brand: G-Shock
	`)

	// using backslashes inside a string:
	fmt.Println(`R:\Machine\Rian`)
	fmt.Println("C:\\Machine\rnt")

	// concatenating strings (+)
	// Go creates a new string because strings are immutable in Go (this is not efficient).
	var s3 = "Don't " + "Stop " + "Learn"
	fmt.Println(s3 + "!") // Don't Stop Learn!

	// getting an element (byte) of a string:
	fmt.Println("Element at index 0:", s3[0]) // 68 >>> ASCII Code
	//  a string is in fact a slice of bytes in Go

	fmt.Printf("%s\n", s3)
	fmt.Printf("%q\n", s3)

	// strings are immutable and can't be changed
	// s3[5] = 'x' // => error: Cannon assign to s3[5].
}
