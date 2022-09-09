package main

import "fmt"

func main() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	yourFriends := make([]string, len(friends))
	copy(yourFriends, friends)

	yourFriends[2] = "Dan" // copy() function create a copy of the slice

	fmt.Println(friends, yourFriends)

}
