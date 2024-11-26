package main

import "fmt"

func main() {
	var colors = []string{"red", "purple", "orange", "green", "silver"}
	fmt.Printf("Type of colors:%T\n", colors)
	fmt.Println("length of colors:", len(colors))

	//! the main difference between both two is the, array should have to define the size implicitely whereas the slice doesn't need to add size implicitely it takes dynamically data.
	//? in both code, the first code is of slice and second code is of array

	names := [6]int{1, 23, 5, 5, 34, 0}
	fmt.Printf("Type of names:%T\n", names)
	fmt.Println("length of names:", len(names))

	// Creating a slice using make() function
	prices := make([]int, 5, 10)
	prices[0] = 11
	prices[1] = 22
	prices[2] = 33
	prices[3] = 44
	prices[4] = 55
	prices = append(prices, 1, 2, 3, 45, 5, 34)
	fmt.Printf("Type of prices:%T\n", prices)
	fmt.Println("length of prices:", len(prices))
	fmt.Println("capacity of prices:", cap(prices))
	fmt.Println("prices are:", prices)

}
