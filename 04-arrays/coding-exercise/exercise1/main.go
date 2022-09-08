package main

import "fmt"

func main() {
	// 1. Using the var keyword declare an array called cities with 2 elements of type string.
	var cities [2]string // {"Sumbawa", "Jakarta"}  Don't initialize the array.
	fmt.Printf("%#v\n", cities) // Print out the cities array and notice the value of its elements

	// 2. Declare an array called grades of type [3]float64 and initialize only the first 2 elements using an array literal.
	var grades = [3]float64{
		0.001,
		33.4,
	}
	fmt.Printf("%#v\n", grades) // Print out the cities array and notice the value of its elements

	// 3. Declare an array called taskDone using the ellipsis operator. 
	taskDone := [...]bool{true, false, false, true} // The elements are of type bool.
	fmt.Printf("%#v\n", taskDone)
	

	// 4. Iterate over the array called cities using the classical for loop syntax and len function and print out element by element.
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	// 5. Iterate over grades using the range keyword and print out element by element.
	for index, value := range grades {
		fmt.Println(index, value)
	}
}