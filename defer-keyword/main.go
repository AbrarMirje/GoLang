package main

import "fmt"

//? defer: defer keyword is used to execute the statements just before the ending of the function

func add(first, second int) int {
	return first + second
}

func main() {
	fmt.Println("Starting of the program")

	result := add(10, 20)
	fmt.Println("Result of the addition is:", result)

	defer fmt.Println("Mid of the program") // this will execute at the end
	fmt.Println("End of the program")
}
