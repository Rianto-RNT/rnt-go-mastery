package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)    // unbuffered channels
	c2 := make(chan int, 3) // buffered channels
	_ = c2

	go func(c chan int) {
		fmt.Println("func gorutine START sending data into the channel")
		c <- 10
		fmt.Println("func gorutine AFTER sending data into the channel")

	}(c1)

	fmt.Println("Main gorutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main gorutine START recieving data")
	d := <-c1
	fmt.Println("main gorutine recieving data:", d)

	time.Sleep(time.Second)

}
