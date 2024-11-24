package main

import "fmt"

func main() {
	age := 23
	name := "Abrar"
	height := 5.67872387

	// fmt.Println("Name:", name, "Age:", age, "Height:", height)
	// string formatter (Printf)
	// fmt.Printf("Name: %s, Age: %d, Height: %f\n", name, age, height)
	fmt.Printf("Name: %s, Age: %d, Height: %.3f\n", name, age, height)

	// To find the type of the variable
	fmt.Printf("Type of age is: %T\n", age)
}
