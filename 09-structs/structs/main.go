package main

import "fmt"

func main() {
	title1, author1, year1 := "The Devine Comedy", "Ryan Morrison", 1815
	title2, author2, year2 := "Macbeth", "Rianto RNT", 1606

	fmt.Println("Book1", title1, author1, year1)
	fmt.Println("Book2", title2, author2, year2)

	type book struct {
		title  string
		author string
		year   int
	}

	type book1 struct {
		title, author string
		year          int
	}

	myBook := book{"The Devine Comedy", "Ryan Morrison", 1815}
	fmt.Println(myBook)

	bestBook := book{title: "Farming", year: 1999, author: "Rianto"}
	_ = bestBook

	aBook := book{title: "random book"}
	fmt.Printf("%#v\n", aBook)
}
