package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://chat.openai.com/"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)

	response.Body.Close() //caller's responsibility to close the connection
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(databytes))
}
