package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to the user input module"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the name for the confirmation : ")

	// comma || error, ok syntax

	// ('\n') mean it will read the input until the \n is received which mean when we press a enter
	input, err := reader.ReadString('\n')
	if err != nil {
	}
	fmt.Println("Thanks for the rating of : ", input)
	fmt.Printf("Type of the rating is %T\n", input)
}
