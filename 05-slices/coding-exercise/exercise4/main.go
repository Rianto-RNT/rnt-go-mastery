// run this program in a terminal with arguments
// ex: go run main.go 5 7.1 9.9 10

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 { //if not run with arguments
		log.Fatal("Please run the program with arguments (2-10 numbers)!")

	}

	//taking the arguments in a new slice
	args := os.Args[1:]

	// declaring and initializing sum and product of type float64
	sum, product := 0., 1.

	if len(args) < 2 || len(args) > 10 {
		fmt.Println("Please enter between 2 and 10 numbers!")
	} else {

		for _, v := range args {
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				continue //if it can't convert string to float64, continue with the next number
			}
			sum += num
			product *= num
		}
	}

	fmt.Printf("Sum: %v, Product: %v\n", sum, product)

}
