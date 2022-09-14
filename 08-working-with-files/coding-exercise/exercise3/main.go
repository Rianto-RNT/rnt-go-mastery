package main

import (
	"log"
	"os"
)

func main() {
	/* 1) Create File*/
	// newFile, err := os.Create("rnt_data.txt")

	// // Error handling
	// if err != nil {
	// 	// log the error and exit the program
	// 	log.Fatal(err) // the idiomaitc way to handle errors

	// }
	// newFile.Close()
	/*End of create file*/

	/*Remove File*/
	err := os.Remove("rnt_data.txt")
	// Error handling
	if err != nil {
		log.Fatal(err)
	}
	/*End of remove file*/

}
