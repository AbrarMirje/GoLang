package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning about url in Go")
	myUrl := "https://jsonplaceholder.typicode.com/todos/1"
	parseUrl, error := url.Parse(myUrl)
	if error != nil {
		fmt.Println("Can't parse URL", error)
	}
	fmt.Printf("Type of url:%T\n", parseUrl)
	fmt.Println("Scheme:", parseUrl.Scheme)
	fmt.Println("Host:", parseUrl.Host)
	fmt.Println("Path:", parseUrl.Path)
	fmt.Println("RawQuery:", parseUrl.RawQuery)

	// modifying the url
	parseUrl.Path = "/golang"

	// Constructing a url string from a url object
	newUrl := parseUrl.String()
	fmt.Println("New URL:", newUrl)
}
