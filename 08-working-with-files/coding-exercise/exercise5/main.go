package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// opening the file in read-only mode. The file must exist (in the current working directory)
	// use a valid path!
	file, err := os.Open("rnt_info.txt")
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	// defer closing the file
	defer file.Close()

	// The file value returnd by osOpen() is wrapped in a bufio.Scanner just like a buffered reader.
	scanner := bufio.NewScanner(file)

	// Reading the wholefile in line by line:
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Checking for any possible errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
