package main

import (
	"log"
	"os"
)

func main() {
	// cheking if files exist
	_, err := os.Stat("rnt_info.txt")
	// Error handling
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file does not exist!")
		}
	}

	// Renaming the file
	err = os.Rename("rnt_info.txt", "rnt_data.txt")
	// Error handling
	if err != nil {
		log.Fatal(err)
	}
}
