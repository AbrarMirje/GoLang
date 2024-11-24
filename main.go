package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learning GoLang with Dedication")

	var name string = "Abrar Mirje"
	fmt.Println(name)

	const PI float64 = 3.147
	// PI = 23.5	// const cannot be modified
	fmt.Println(PI)

	personAge := 23
	fmt.Println(personAge)

	// PublicVariable can exported to the another packages bcoz of the initial letter is starting with capital letter
	// var PublicVariable int = 32
	// privateVariable cannot be exported to the another packages bcoz of the initial letter is starting with smaller letter
	// var privateVariable int = 99

}
