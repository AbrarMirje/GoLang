package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 1. Creation of file
	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("Error while creating a file")
	// 	return
	// }
	// defer file.Close()

	// 2. Writing into the file
	// content := "Hello World from Abrar"
	// first way to write
	// file.WriteString(content)

	// second way to write
	// io.WriteString(file, content)

	// fmt.Println("File created successfully")

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error occured while creating the file")
		return
	}
	defer file.Close()

	// create a buffer to read the file content
	// buffer is a temperory memory storage
	buffer := make([]byte, 1024)

	// Read the file content into the buffer
	for {
		n, errors := file.Read(buffer)
		if errors == io.EOF {
			break
		}
		if errors != nil {
			fmt.Println("Error while reading the file")
			return
		}

		// Process the read content
		fmt.Println(string(buffer[:n]))
	}
}
