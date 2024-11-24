package main

import (
	"bufio"
	"fmt"
	"os"
)

func simpleFunction() {
	fmt.Println("This is just a simple function")
}

// Function with parameters

func add(firstNumber, secondNumber int) int {
	return firstNumber + secondNumber
}

func mult(firstNumber, secondNumber int) (result int) {
	result = firstNumber * secondNumber
	return
}
func div(firstNumber, secondNumber int) (result int) {
	return firstNumber / secondNumber

}

func greetings(name string, age int) string {
	return fmt.Sprintf("Name: %s, Age: %d", name, age)
}

// ! This code is not printing anything, we will solve that issue very soon.
func gettingUserInputAndReturningIt() {

	fmt.Println("What's your name?")
	var input string
	fmt.Scan(&input)
	reader := bufio.NewReader(os.Stdin)
	output, _ := reader.ReadString('\n')
	fmt.Println(output)
}

func main() {
	fmt.Println("We are learning functions in GoLang")
	simpleFunction()
	result := add(10, 20)
	fmt.Println(result)

	greet := greetings("Stephen", 23)
	fmt.Println(greet)

	gettingUserInputAndReturningIt()

	multiplication := mult(5, 6)
	fmt.Println("Multiplication is:", multiplication)

}
