package main

import "fmt"

func main() {
	// 1. Using the var keyword declare a string called name and initialize it with your name
	var name string = "Ryan"
	fmt.Println(name)
	// 	2. Using short declaration syntax declare a string called country and assign the country you are living in to the string variable.
	country := "Indonesia"
	fmt.Println(country)
	fmt.Printf("Name: %s\nCountry: %s\n", name, country)

	// 3. Print the following string on multiple lines like this:
	// Your name: `here the value of name variable`
	// Country: `here the value of country variable`
	//equivalent to:
	fmt.Printf(`Name: %s
Country: %s
`, name, country)

	// 4. Print out the following strings:
	// a) He says: "Hello"
	fmt.Println(`Ryan Morrison says: "Hello"`)
	fmt.Println("C:\\Users\\rnt.txt")

	// b) C:\Users\a.txt
}
