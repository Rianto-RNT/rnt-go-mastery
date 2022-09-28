package main

import (
	"fmt"
	"sync"
)

func sum(a, b float64, wg *sync.WaitGroup) {
	s := a + b
	fmt.Printf("sum of %.2f and %.2f is %.2f\n", a, b, s)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go sum(10.3, 20.1, &wg)
	go sum(5, 7, &wg)
	go sum(33.5, 33.66, &wg)

	wg.Wait()
}
