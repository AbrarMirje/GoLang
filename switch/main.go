package main

import "fmt"

func main() {
	// 1. Basic switch statement
	day := 4
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	default:
		fmt.Println("Invalid day")
	}

	// 2. switch statement with multiple values
	month := "Abrar"
	switch month {
	case "January", "February", "March":
		fmt.Println("Winter")
	case "April", "May", "June":
		fmt.Println("Spring")
	case "July", "August", "September":
		fmt.Println("Summer")
	case "October", "November", "December":
		fmt.Println("Autumn")
	default:
		fmt.Println("Invalid month")
	}

	// 3. switch with expression
	temperature := 25
	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature > 0 && temperature < 10:
		fmt.Println("Cold")
	case temperature >= 10 && temperature < 20:
		fmt.Println("Cool")
	case temperature >= 20 && temperature < 30:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}

}
