package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning web api")

	response, error := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if error != nil {
		fmt.Println("Error while returning response", error)
		return
	}
	defer response.Body.Close()

	// Reading the response body
	data, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println("Error while reading response body", error)
		return
	}
	fmt.Println("response:", string(data))
}
