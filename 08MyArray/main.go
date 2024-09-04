package main

import "fmt"

func main() {

	fmt.Println("Welcome to the Go lang of Array concept")

	var nameList [4]string

	nameList[0] = "Surya"
	nameList[1] = "Selva"
	nameList[2] = "Mohan"
	nameList[3] = "Karthi"
	fmt.Println(len(nameList))
	fmt.Println(nameList)

	var names = [4]string{"Sibi", "Vimal", "Sakthi", "Chandru"}
	fmt.Println(names)
}
