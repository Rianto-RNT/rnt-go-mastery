package main

import (
	"fmt"
	"sync"
)

func sayHello(n string, w *sync.WaitGroup) {
	fmt.Printf("Hello, %s!\n", n)
	w.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go sayHello("Mr. Ryan", &wg)

	wg.Wait()
}
