package main

import "fmt"

func main() {
	age := 41
	// if age >= 18 {
	// 	fmt.Print("You are eligible for voting")
	// } else {
	// 	fmt.Print("You still have to grow up!")
	// }

	if age >= 18 && age <= 40 {
		fmt.Print("You are a young man")
	} else if age > 40 {
		fmt.Print("You are becoming a old one")
	} else {
		fmt.Print("You are too young")
	}
}
