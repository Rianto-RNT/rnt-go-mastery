package main

import "fmt"

func sayHello(n string) {
	fmt.Printf("Hello, %s!\n", n)

}

func main() {
	go sayHello("Mr. RNT")
}
