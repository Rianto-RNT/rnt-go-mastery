package main

import "fmt"

func main() {
	var ch chan int
	fmt.Println(ch)

	ch = make(chan int)
	fmt.Println("Valueof ch:", ch)

	c := make(chan int)

	// <- channel operator

	// SEND
	// c <- 10

	// RECIEVE
	num := <-c

	fmt.Println(<-c)

	_ = num

	close(c)
}
