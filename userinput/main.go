package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What's your name?")
	// var username string
	// fmt.Scan(&username)
	// but using below line we only read the single string like if we have Abrar Mirje then it only read the Abrar but not Mirje, bcoz there is a whitespace in between Abrar Mirje so that Scan() assumes that we have to break the line so that it breaks the line from there and that's the reason we cannot able to read the string after the whitespace
	// fmt.Println("Hello,", username)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println(name)

}
