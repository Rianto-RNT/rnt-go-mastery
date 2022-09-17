package main

import "fmt"

func f1(n uint) (int, int) {

	fact := 1
	sum := 0

	for i := 1; i <= int(n); i++ {
		fact *= i
		sum += i
	}
	return fact, sum

}

func main() {
	f, s := f1(4)
	fmt.Println("f:", f, "s:", s)

	f, s = f1(10)
	fmt.Println("f:", f, "s:", s)
}
