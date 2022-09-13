package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var newFile *os.File
	fmt.Printf("%T\n", newFile)

	var err error

	newFile, err = os.Create("r.txt")
	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		log.Fatal(err)
	}

	err = os.Truncate("r.txt", 0)
	if err != nil {
		log.Fatal(err)
	}
	newFile.Close()

	file, err := os.OpenFile("r.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("r.txt")
	p := fmt.Println

	p("File Name:", fileInfo.Name())
	p("Size in bytes:", fileInfo.Size())
	p("Last modified:", fileInfo.ModTime())
	p("Is directory?:", fileInfo.IsDir())
	p("Permissions:", fileInfo.Mode())

	fileInfo, err = os.Stat("b.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File file does not exist!")
		}
	}

	// oldPath := "r.txt"
	// newPath := "rrr.txt"
	// err = os.Rename(oldPath, newPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = os.Remove("rrr.txt")
	if err != nil {
		log.Fatal(err)
	}
}
