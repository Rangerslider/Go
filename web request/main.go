package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev" //web link

func main() {
	fmt.Println("Web request")

	response, err := http.Get(url) // The HTTP GET request method is used to request a resource from the server.

	if err != nil {
		panic(err) // catch the error and stop code compilation if any
	}

	fmt.Printf("Response is of type: %T\n", response) // response has http type

	defer response.Body.Close() // caller's responsibility to close the connection .. ones u request then you need to close it

	databytes, err := ioutil.ReadAll(response.Body) // reading response // here read all with body

	if err != nil {
		panic(err)
	}
	content := string(databytes) // databytes to readable string conversion
	fmt.Println(content)
}
