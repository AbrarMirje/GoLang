package main

import "fmt"

func main() {
	fmt.Println("We are now diving into the arrays in GoLang")

	// Creation of an array
	var names [5]string
	names[0] = "Abrar"
	names[3] = "Mirje"
	names[4] = "Victus"
	fmt.Println("Names are:", names)

	// arrays initialization and declaration at same line
	var colors = [3]string{"Red", "Green", "Blue"}
	fmt.Println("Colors are:", colors)

	// Another way to initialize and declare arrays at same line
	var fruits = [...]string{"Apple", "Banana", "Cherry", "Date", "Watermelon"}
	fmt.Println("Fruits are:", fruits)

	// len() function to find the length of the array
	fmt.Println("Length of the array is:", len(fruits))

	// default values of array elements according to their types
	var numbers [5]int
	fmt.Println("Numbers are:", numbers) // [0 0 0 0 0]

	var sports [3]string                       // have a single space for each element
	fmt.Println("Sports are:", sports)         // [   ]
	fmt.Printf("Price array is: %q\n", sports) // ["" "" ""]

	var isAvailable [2]bool
	fmt.Println("Is available:", isAvailable) // [false false]

	var prices [3]float64
	fmt.Println("Prices are:", prices) // [0 0 0]

	var person [5]string
	person[0] = "Abrar"
	person[1] = "Mirje"
	person[2] = "Victus"
	fmt.Println("Length:", len(person))

}
