package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "https://chaicode.com/"

func main() {

	fmt.Println("Let's learn some net/http")

	response, err := http.Get(URL)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%T\n", response)

	defer response.Body.Close() // Our responsibility to close the connection

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(databytes))
}
