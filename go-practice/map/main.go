package main

import "fmt"

func main() {
	// Storing age of the students
	studentsAge := make(map[string]int)
	studentsAge["John"] = 20
	studentsAge["Alice"] = 22
	studentsAge["Bob"] = 21
	fmt.Printf("Age of Bob:%d\n", studentsAge["Bob"])

	studentsAge["Bob"] = 54
	fmt.Printf("Age of Bob:%d\n", studentsAge["Bob"])

	delete(studentsAge, "Bob")
	fmt.Printf("Age of Bob:%d\n", studentsAge["Bob"])

	// checking is there key present or not
	key, isPresent := studentsAge["Alice"]
	if isPresent {
		fmt.Printf("Age of Alice:%d\n", key)
	} else {
		fmt.Println("Alice is not present in the class")
	}

	//! Another way of creation of map
	studentAge := map[string]int{
		"John":  20,
		"Alice": 22,
		"Bob":   21,
	}

	// we can iterate over the map using for loop
	for key, value := range studentAge {
		fmt.Printf("Key:%s Value:%d\n", key, value)
	}

}
