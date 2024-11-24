package main

import "fmt"

func divide(firstNumber, secondNumber float64) (float64, error) {
	if secondNumber == 0 {
		return 0, fmt.Errorf("denominator must be zero")
	}
	return firstNumber / secondNumber, nil
}

func main() {
	// fmt.Println("Error handling in the functions")
	// fmt.Println("Division is:", divide(13, 0))

	// result, _ := divide(10, 0)
	// fmt.Println("Division is:", result)

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error Handling")
	}
	fmt.Println("Division is:", result)

	//? _ (underscore) is used when function is returning something but we have to ignore that thing so that time we will use _ (underscore) symbol
}
