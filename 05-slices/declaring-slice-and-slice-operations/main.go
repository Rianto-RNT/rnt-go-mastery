package main

import "fmt"

func main() {
	var cities []string
	fmt.Println("Cities is equal to nill:", cities == nil)
	fmt.Printf("Cities %#v\n", cities)
	fmt.Println(len(cities))

	numbers := []int{2, 3, 4, 5}
	fmt.Println(numbers)

	nums := make([]int, 2)
	fmt.Printf("%#v\n", nums)

	type names []string
	friends := names{"Rian", "RNT"}
	fmt.Println(friends)

	myFriend := friends[0]
	fmt.Println("Friend is", myFriend)

	friends[0] = "Saprol"
	fmt.Println("My best friend is", friends[0])

	for index, value := range numbers {
		fmt.Printf("Index: %v value: %v\n", index, value)
	}

	var n []int
	n = numbers
	fmt.Println(n)
}
