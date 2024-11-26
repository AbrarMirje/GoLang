package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Calling store api")
	response, error := http.Get("https://fakestoreapi.com/products/3")
	if error != nil {
		fmt.Println("Error calling store api")
		return
	}
	defer response.Body.Close()

	data, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println("Error reading response")
		return
	}
	fmt.Println("storeData:", string(data))
}
