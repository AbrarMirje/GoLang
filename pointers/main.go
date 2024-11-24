package main

import "fmt"

//? Pointer: Pointer stores the memeory address of the other variable. It provides way to work with memory directly

func modifyValueByReference(num *int) {
	*num = *num * 2
}

func main() {
	// 1. First method
	var number int
	number = 10

	var ptr *int
	ptr = &number

	fmt.Println("Number is:", number)
	fmt.Println("Address is:", ptr)

	// 2. Second method
	name := "GoLang"
	pointer := &name
	fmt.Println("Value is:", pointer)
	fmt.Println("Value contains data is:", *pointer)

	// Create a function to modify the data
	value := 30
	modifyValueByReference(&value)
	fmt.Println("Value contains:", value)
}
