package main

import "fmt"

func print(msg string) {
	fmt.Println(msg)
}

func main() {
	defer print("Whatever im late anyway")
	fmt.Println("Good Afternoon, Rianto.")
}
