package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	House   int
	Area    string
	State   string
	Country string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	var abrar Person
	fmt.Println("Abrar person:", abrar) // Abrar person: {  0}
	abrar.FirstName = "Abrar"
	abrar.LastName = "Mirje"
	abrar.Age = 23
	fmt.Println("Abrar person:", abrar) // Abrar person: {  0}

	// Another way to declare the struct
	person1 := Person{
		FirstName: "Robert",
		LastName:  "Downey Jr.",
		Age:       55,
	}

	fmt.Println("Robert person:", person1) // Robert person: {Robert Downey Jr.

	// by using new keyword
	// new keyword is used to allocate memory for a struct
	// it returns a pointer to the struct
	// so we need to dereference it using *
	person2 := new(Person)
	person2.FirstName = "Alex"
	person2.LastName = "Petry"
	person2.Age = 32
	fmt.Println("Alex person:", person2)  // Alex person: &{Alex Petry 32}
	fmt.Println("Alex person:", *person2) // Alex person: {Alex Petry 32

	employee1 := Employee{
		Person_Details: Person{
			FirstName: "John",
			LastName:  "Doe",
			Age:       30,
		},
		Person_Contact: Contact{
			Phone: "1234567890",
			Email: "john.doe@example.com",
		},
		Person_Address: Address{
			House: 123,
			Area:  "New York",
			State: "NY",
		},
	}
	fmt.Println("-----------------------")
	fmt.Println(employee1)
	fmt.Println(employee1.Person_Address.State)

	// adding dynamic values in struct
	employee1.Person_Address.Country = "US"
	fmt.Println(employee1.Person_Address.Country)
}
