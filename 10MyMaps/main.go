package main

import "fmt"

func main() {

	fmt.Println("Let's start our map of Go ")

	languages := make(map[string]string)

	languages["Go"] = "Go"
	languages["Jv"] = "Java"
	languages["Py"] = "Python"
	languages["Rb"] = "Ruby"
	languages["Js"] = "JavaScript"

	fmt.Println("Iterate through the key and value of languages")

	for k, v := range languages {
		fmt.Println("Key:", k, "Value:", v)
	}
}
