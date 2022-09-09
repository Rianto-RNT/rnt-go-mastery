package main

import "fmt"

func main() {
	/* ERROR CODE TO FIX >>>
	mySlice := []float64{1.2, 5.6}

	mySlice[2] = 6

	a := 10
	mySlice[0] = a

	mySlice[3] = 10.10

	mySlice = append(mySlice, a)

	fmt.Println(mySlice)
	*/

	mySlice := []float64{1.2, 5.6}

	// ERROR -> index out of range [2] with length 2
	// mySlice[2] = 6
	mySlice[1] = 6

	// FINAL RESULT >>>
	// ERROR -> cannot use a (type int) as type float64 in assignment
	// a := 10
	a := 10.
	mySlice[0] = a

	// ERROR -> index out of range [3] with length 2
	// mySlice[3] = 10.10
	mySlice[1] = 10.10

	mySlice = append(mySlice, a)

	fmt.Println(mySlice)

}
