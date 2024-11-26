package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getRequest() {
	response, error := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if error != nil {
		fmt.Println(error)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error in getting response", response.Status)
		return
	}

	// data, error := ioutil.ReadAll(response.Body)
	// if error != nil {
	// 	fmt.Println(error)
	// 	return
	// }
	// fmt.Println("Data:", string(data))

	//! both code are same but the Decode one is more efficient and used.

	//? Recommended
	var todo Todo
	er := json.NewDecoder(response.Body).Decode(&todo)
	if er != nil {
		fmt.Println("Error while converting the data:", er)
		return
	}
	fmt.Println("Todo Data:", todo)
}

func postRequest() {
	todo := Todo{
		UserId:    123,
		Title:     "Buy Milk",
		Completed: false,
	}
	// Converting to JSON (Marshalling)
	jsonData, error := json.Marshal(todo)
	if error != nil {
		fmt.Println(error)
		return
	}
	// fmt.Println("jsonData:", string(jsonData))
	jsonString := string(jsonData)

	// convert string to jsonReader
	jsonReader := strings.NewReader(jsonString)
	myUrl := "https://jsonplaceholder.typicode.com/todos"

	resp, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error while posting the data:", err)
		return
	}
	defer resp.Body.Close()

	readableData, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("data:", string(readableData))
	fmt.Println("Response:", resp.Status)

}

func putRequest() {
	todo := Todo{
		UserId:    2563,
		Title:     "Buy Books",
		Completed: true,
	}
	jsonData, error := json.Marshal(todo)
	if error != nil {
		fmt.Println("Error while marshalling")
		return
	}

	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	// Convert JSON to string
	jsonFormat := string(jsonData)

	// Convert string to reader
	jsonReader := strings.NewReader(jsonFormat)

	// Creat PUT request
	request, error := http.NewRequest(http.MethodPut, myUrl, jsonReader)
	if error != nil {
		fmt.Println("Error while updating:", error)
		return
	}
	request.Header.Set("Content-type", "application/json")

	// Create a client and send request
	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("Error while sending request", err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Data:", string(data))
	fmt.Println("Status Code:", res.Status)
}

func deleteRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	// Creat DELETE request
	request, error := http.NewRequest(http.MethodDelete, myUrl, nil)
	if error != nil {
		fmt.Println("Error while deleting:", error)
		return
	}

	// Create a client and send request
	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("Error while sending request", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Status Code:", res.Status)

}

func main() {
	fmt.Println("Data is loading...")
	// getRequest()
	// postRequest()
	// putRequest()
	deleteRequest()
}
