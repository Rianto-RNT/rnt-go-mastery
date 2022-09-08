package main

import "fmt"

func main() {
    /* ERROR 
	myArray := [3]float64{1.2, 5.6}

    myArray[2] = 6

    a := 10
    myArray[0] = a

    myArray[3] = 10.10

    fmt.Println(myArray)
	*/

	// RUN ME 
	myArray := [3]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10
	// ERROR -> invalid array index 3 (out of bounds for 3-element array)
	// myArray[0] = a
	myArray[0] = float64(a)

	// ERROR -> invalid array index 3 (out of bounds for 3-element array)
	// myArray[3] = 10.10

	myArray[2] = 10.10

	fmt.Println(myArray)


}