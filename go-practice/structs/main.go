package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	Street  string
	City    string
	Country string
}

type Employee struct {
	Person_Details  Person
	Contact_Details Contact
	Work_Address    Address
}

func main() {
	var steve Person
	steve.Name = "Steve"
	steve.Age = 30
	fmt.Println(steve)

	david := Employee{
		Person_Details: Person{
			Name: "David",
			Age:  25,
		},
		Contact_Details: Contact{
			Email: "david@example.com",
			Phone: "1234567890",
		},
		Work_Address: Address{
			Street:  "123 Main St",
			City:    "Anytown",
			Country: "USA",
		},
	}

	fmt.Println(david)
	fmt.Printf("Email of david:%s\n", david.Contact_Details.Email)
}
