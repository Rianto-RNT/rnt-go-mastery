package main

import "fmt"

func main() {
	// Consider the following array declaration: nums := []int{30, -1, -6, 90, -6}
	nums := []int{30, -1, -6, 90, -6}
	
	// Write a Go program that counts the number of positive even numbers in the array.
	var count int

	for _, v := range nums {
		if v%2 == 0 && v > 0 {
			count++
		}
	}

	fmt.Printf("%#v\n",nums)
	fmt.Println("No. of positive even numbers in nums: ", count)
}