package main

import (
	"log"
	"os"
)

func main() {
	newFile, err := os.Create("rnt_exercise1.txt")

	// Error handling
	if err != nil {
		// log the error and exit the program
		log.Fatal(err) // the idiomaitc way to handle errors

	}
	newFile.Close()
}
