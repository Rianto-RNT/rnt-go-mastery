package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Openig the file read-only mode. The file must exist (in the current working directory)
	// Use valid path
	file, err := os.Open("rnt_info.txt")
	// Error handling
	if err != nil {
		log.Fatal(err)
	}

	// Defer closing the file
	defer file.Close()

	// ioutil.ReadAll() reads from the file until an error of EOF and return the data it read
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as a string: %s\n", data)
}
