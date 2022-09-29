package main

import "fmt"

func main() {
	ch := make(chan string)

	go func(n string) {
		ch <- n
	}("RNT!")

	value := <-ch
	fmt.Println("Value recieved from channel:", value)
}
