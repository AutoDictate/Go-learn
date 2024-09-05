package main

import "fmt"

func main() {

	defer fmt.Println("Hello World from defer")
	fmt.Println("Hello World ")
	myDefer()
}

func myDefer() {

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
