package main

import (
	"fmt"
	"sync"
)

func deposit(b *int, n int, wg *sync.WaitGroup) {
	*b += n
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup) {
	*b -= n
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg)
		go withdraw(&balance, i, &wg)
	}
	wg.Wait()

	fmt.Println("Final balance value:", balance)
}
