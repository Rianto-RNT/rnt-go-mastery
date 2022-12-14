package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(50)

	// Important: i start fom 100. not 100
	// i is float64, not int. math.Sqrt() takes in a float46
	for i := 100.; i < 150.; i++ {
		go func(f float64, wg *sync.WaitGroup) {
			x := math.Sqrt(f)
			fmt.Printf("Square root of %.2f is %.6f\n", f, x)
			wg.Done()
		}(i, &wg)

		wg.Wait()
	}
}
