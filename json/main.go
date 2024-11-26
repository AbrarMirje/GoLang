package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	person1 := Person{Name: "Stephen", Age: 32, IsAdult: true}
	fmt.Println("stephen:", person1)

	// converting person data into JSON encoding (Marshaling)
	data, error := json.Marshal(person1)
	if error != nil {
		fmt.Println("Error while marshalling: ", error)
		return
	}
	fmt.Println("Person data is: ", string(data))

	// Decoding (Unmarshalling)
	var personData Person
	err := json.Unmarshal(data, &personData)
	if err != nil {
		fmt.Println("Erro while unmarshalling")
		return
	}
	fmt.Println("Person data: ", personData)

}
