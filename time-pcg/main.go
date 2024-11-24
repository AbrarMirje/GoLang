package main

import (
	"fmt"
	"time"
)

func main() {
	// Go represents time in yyyy-mm-dd hh:mm:ss format and that is the 2006-01-02 15:04:05

	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Printf("Type of current time %T\n", currentTime) // time.Time

	// If we have to format the time then
	formattedTime := currentTime.Format("02-01-2006, Monday, 15:04:05")
	fmt.Println(formattedTime)

	// If we want AM or PM
	formatAMorPM := currentTime.Format("02-01-2006, 3:04 PM")
	fmt.Println(formatAMorPM)

	//? string data converting into time format
	layout_str := "2006-01-02"
	dateStr := "24-11-2024"
	formatted_time, _ := time.Parse(layout_str, dateStr)
	fmt.Println(formatted_time)
	
}
