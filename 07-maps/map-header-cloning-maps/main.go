package main

import "fmt"

func main() {
	friends := map[string]int{"Rian": 30, "saprol": 27}

	neighbors := friends

	friends["Rian"] = 17

	fmt.Println(neighbors)

	people := make(map[string]int)

	for k, v := range friends {
		people[k] = v
	}

	people["rnt"] = 22

	fmt.Printf("%#v >>> %#v\n", people, friends)

	delete(friends, "Rian")
	fmt.Println(friends)

	delete(people, "Morrison")
}
