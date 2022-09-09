package main

import "fmt"

func main() {
	friends := []string{"Rian", "Rnt", "Saprol", "Udin"}
	yourFriends := []string{}

	yourFriends = append(yourFriends, friends...)

	// yourFriends[4] = "Morrison" // panic: runtime error: index out of range [4] with length 4
	yourFriends[3] = "Morrison"

	fmt.Println(friends, yourFriends)
}
