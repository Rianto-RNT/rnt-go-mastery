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

	// page := lastBook.pages

	fmt.Printf("%#v\n", lastBook)

	lastBook.author = "Ryan Morrison"
	lastBook.year = 1815
	fmt.Printf("%#v\n", lastBook)

	aBook := book{title: "Programming with Go", author: "Ryan Morrison", year: 1815}
	fmt.Println(aBook == lastBook)

	myBook := aBook
	myBook.year = 2022
	fmt.Println(myBook, aBook)
}
