package main

import (
	"fmt"
	"strings"
)

func searchItem(mySlice []string, myStr string) bool {
	for _, v := range mySlice {
		if strings.EqualFold(v, myStr) {
			return true
		}
	}
	return false
}

func main() {
	animals := []string{"Troak", "Tanenang", "Kleang"}

	result := searchItem(animals, "KLEANG")
	fmt.Println(result) // >>> true

	result = searchItem(animals, "TaNenAnG")
	fmt.Println(result) // >>> true

	result = searchItem(animals, "bawi")
	fmt.Println(result) // >>> false
}
