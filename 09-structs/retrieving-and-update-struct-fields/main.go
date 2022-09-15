package main

import "fmt"

func main() {
	type book struct {
		title  string
		author string
		year   int
	}

	lastBook := book{title: "Programming with Go"}
	fmt.Println(lastBook.title)
}
