package main

import "fmt"

func searchItem(mySlice []string, myStr string) bool {
	for _, v := range mySlice {
		if v == myStr {
			return true
		}
	}
	return false
}

func main() {
	animals := []string{"Troak", "Tanenang", "Kleang"}

	result := searchItem(animals, "Kleang")
	fmt.Println(result) // >>> true

	result = searchItem(animals, "Jaran")
	fmt.Println(result) // >>> false
}
