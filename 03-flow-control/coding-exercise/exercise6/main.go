package main

import "fmt"

func main() {
	// IF STATEMENT
	// age := -9
	// if age < 0 || age > 100 {
	// 	fmt.Println("Invalid Age")
	// } else if age <= 18 {
	// 	fmt.Println("You are minor!")
	// } else if age == 18 {
	// 	fmt.Println("Congratulations! You've just become major!")
	// } else {
	// 	fmt.Println("You are major!")
	// }

	// SWITCH STATEMENTS
	age := 10
	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age <= 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}
}
