package main

import "fmt"

func main() {
	s1 := "Hi there, RNT!"
	fmt.Println(s1)

	fmt.Println("Good Morning: \"RNT\"")

	fmt.Println(`Good Afternoon: "Rianto"`)

	s2 := `Whatever im late anyway`
	fmt.Println(s2)

	fmt.Println("Price: 150\nBrand: G-Shock")

	fmt.Println(`
Price: 150
Brand: G-Shock
	`)

	fmt.Println(`R:\Machine\Rian`)
	fmt.Println("C:\\Machine\rnt")

	var s3 = "Don't " + "Stop " + "Learn"
	fmt.Println(s3 + "!")

	fmt.Println("Element at index 0:", s3[0])

	// S3[5] = 'x'

	fmt.Printf("%s\n", s3)
	fmt.Printf("%q\n", s3)
}
