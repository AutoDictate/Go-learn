package main

import "fmt"

func main() {

	fmt.Println("Let's start our map of Go ")

	languages := make(map[string]string)

	name := "surya"
	names := make(map[string]int)

	t := string(name[1])
	names[t] = 23

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
