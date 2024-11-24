package main

import "fmt"

func main() {

	// mapping students grades with their names
	// it maps in key value pairs
	// it is unorder collection of data
	studentsGrades := make(map[string]int)
	studentsGrades["John"] = 90
	studentsGrades["Alice"] = 85
	studentsGrades["Bob"] = 95
	studentsGrades["Charlie"] = 80

	// retrive the marks of Alice
	fmt.Printf("Grades of Alice: %d\n", studentsGrades["Alice"])

	// reassign of the marks for the Alice
	studentsGrades["Alice"] = 92
	fmt.Printf("New Grades of Alice: %d\n", studentsGrades["Alice"])

	// deleting the student
	delete(studentsGrades, "Alice")
	fmt.Printf("New Grades of Alice: %d\n", studentsGrades["Alice"])

	// checking if a key exists or not
	//? map returns two values, first is the value of the key and second is if key is exists or not, if key is exists then it returns true else returns false
	grades, exists := studentsGrades["Abrar"]
	if exists {
		fmt.Printf("Grades of Abrar: %d\n", grades)
	} else {
		fmt.Printf("Abrar is not in the map\n")
	}

	// using range function
	// it returns the key and value of the map
	// it is used to iterate over the map
	for key, value := range studentsGrades {
		fmt.Printf("Student: %s, Grades: %d\n", key, value)
	}

	//! another way of creation of an map
	//? it is used when we know the keys and values at the time of creation of map
	person := map[string]int{
		"John":    25,
		"Alice":   22,
		"Bob":     28,
		"Charlie": 20,
	}
	fmt.Println("------------------------------------")
	for key, value := range person {
		fmt.Printf("Person: %s, Age: %d\n", key, value)
	}
}
