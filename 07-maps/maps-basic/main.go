package main

import "fmt"

func main() {
	var employees map[string]string
	fmt.Printf("%#v\n", employees)

	fmt.Printf("No of pairs: %d\n", len(employees))

	fmt.Printf("The value for key %q is %q", "Saprol", employees["Rian"])

	var accounts map[string]float64
	fmt.Printf("%#v\n", accounts["0x323"])

	var m1 map[[5]int]string
	_ = m1

	// employees["Rian"] = "Man"

	people := map[string]float64{}

	people["John"] = 21.4
	people["Marry"] = 10
	fmt.Println(people)

	map1 := make(map[int]int)
	map1[4] = 8

	balances := map[string]float64{
		"USD": 323.11,
		"Rp":  25.000,
		// 50: 322.1, // Error, all keys are the same type
		"CHF": 21258.1,
	}

	fmt.Println(balances)

	m := map[int]int{1: 10, 2: 20, 3: 30}
	_ = m

	balances["USD"] = 500.2
	balances["Rp"] = 30.000
	fmt.Println(balances)

	fmt.Println(balances["RON"])

	balances["RON"] = 0

	v, ok := balances["RON"]

	if ok {
		fmt.Println("The RON balances is: ", v)
	} else {
		fmt.Println("The RON keys doesn't exist in the map")
	}

	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)

		delete(balances, "USD")
		fmt.Println(balances)
	}
}
