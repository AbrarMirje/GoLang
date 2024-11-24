package main

import (
	"fmt"
	"strings"
)

func main() {
	//? 1) Splitting the number on the basis of the comma
	data := "apple, orange, banana, watermelon"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	//? 2) Counting the occurence of the number
	str := "one two three four two two five"
	counter := strings.Count(str, "two")
	fmt.Println("Counter:", counter)

	//? 3) Trimmed white space
	white_space := "                     Hello Abrar           "
	trimmed := strings.TrimSpace(white_space)
	fmt.Println(trimmed)
}
