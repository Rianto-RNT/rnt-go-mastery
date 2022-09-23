package main

import "fmt"

type book struct {
    title string
    price float64
}

// method for book type
/**
the method should receive a pointer to book, not a value
in Go everyhing is passed by value so functions work on copies!
*/
func (b *book) changePrice(p float64) { // NOTE: change price with *
    b.price = p
}

func main() {
    // book value
    bestBook := book{title: "The Trial by Kafka", price: 9.9}

    // changing the price
    bestBook.changePrice(11.99)

    fmt.Printf("Book's price:%.2f\n", bestBook.price) // the price is changed
}