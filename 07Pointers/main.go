package main

import "fmt"

func main() {

	fmt.Println("Welcome to Go Learn of POINTERS concept")

	myNumber := 23

	var ptr *int // ptr is a pointer which is declared here

	ptr = &myNumber // ptr is initialized here

	fmt.Println("Value of the ptr is : ", ptr)
	fmt.Println("Value of the ptr is : ", *ptr)

	*ptr = *ptr + 23

	fmt.Println("New Value of the ptr is : ", *ptr)

}
