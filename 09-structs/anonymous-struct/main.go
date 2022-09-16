package main

import "fmt"

func main() {
	rian := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Rian",
		lastName:  "RNT",
		age:       29,
	}

	fmt.Printf("%#v\n", rian)
	fmt.Printf("Rian's Age: %d\n", rian.age)

	type Book struct {
		string
		float64
		bool
	}

	b1 := Book{"1989, by Ryan Morrison", 11.2, false}
	fmt.Printf("%#v\n", b1)

	fmt.Println(b1.string)

	type Employee struct {
		name   string
		salary int
		bool
	}

	e := Employee{"Rian", 50000, false}
	fmt.Printf("%#v\n", e)
	e.bool = true
	fmt.Printf("%#v\n", e)
}
