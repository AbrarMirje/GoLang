package main

import (
	"fmt"
	"strconv"
)

func main() {
	//? 1) int to float64 conversion
	var num int = 40
	fmt.Println("Number is:", num)
	fmt.Printf("Type of num is:%T\n", num)

	var data float64 = float64(num)
	fmt.Println("Data is", data)
	fmt.Printf("Data type is:%T\n", data)

	//? 2) int to string conversion
	str := strconv.Itoa(num)
	fmt.Println("String is:", str)
	fmt.Printf("str data type is:%T\n", str)

	//? 3) string to int conversion
	number_string := "1234"
	number_int, _ := strconv.Atoi(number_string)
	fmt.Println("number_string is:", number_string)
	fmt.Printf("Type of number_string is:%T\n", number_int)

	//? 4)  string to float conversion
	string_number := "569.25"
	float_number, _ := strconv.ParseFloat(string_number, 64)
	fmt.Println("string_number is:", string_number)
	fmt.Printf("Type of string_number is:%T\n", float_number)

}
