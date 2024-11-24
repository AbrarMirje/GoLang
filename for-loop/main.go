package main

import "fmt"

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// Infinite loop in Go
	// counter := 0
	// for {
	// 	fmt.Println("Infinite loop")
	// 	counter++

	// 	if counter == 5 {
	// 		break
	// 	}
	// }

	//? range funtion in for loop using Go
	// numbers := []int{11, 22, 33, 44, 55}
	// for index, value := range numbers {
	// 	fmt.Printf("Index: %d, Value: %d\n", index, value)
	// }

	data := "Abrar Mirje"
	for index, char := range data {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
}
