package main

import "fmt"

func main() {

	// 1. There is no inheritance in Structs
	// 2. There is no "super" keyword
	// 3. There is parent-child relationship

	fmt.Println("Welcome to Go Learn of topic Structs")

	// Create a new user

	surya := User{
		Name:   "Surya",
		Age:    22,
		Email:  "surya@gmail.com",
		Status: true,
	}
	fmt.Printf("Surya details are : %+v\n", surya)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
