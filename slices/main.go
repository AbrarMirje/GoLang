package main

import "fmt"

// ? slices are used to store data unlike arrays, but we can store data dynamically in slice. So that is the only difference between array and slice. At the time of development we mostly use slices over the array.

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	//! We can append/add data into the current slice
	numbers = append(numbers, 6, 7, 8, 9, 10, 11, 12, 13)

	fmt.Println("Numbers after append:", numbers)

	//? BTS slices are implementing the array but in dynamic form
	fmt.Printf("The type of numbers:%T\n", numbers)
	fmt.Println("Length:", len(numbers))

	// Initializing an empty slice
	name := []string{}
	fmt.Println("Name:", name)
	fmt.Printf("Type of name:%T\n", name)
	fmt.Println("Size of name", len(name))

	// We can create slice using make function
	prices := make([]int, 5, 10)
	prices = append(prices, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println("Prices:", prices)
	fmt.Println("Length:", len(prices))

	// As long as length crosses the capacity, then automatically capacity doubles itself like in this example as long as length increase then the initial capacity increase from 10 to 20
	fmt.Println("Length:", cap(prices))
}
