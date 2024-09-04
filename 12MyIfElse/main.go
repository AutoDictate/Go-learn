package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Learn Go with if-else conditions")
	fmt.Println("Voting Eligibility Application")
	fmt.Print("What is your age? : ")

	var age int

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	age, _ = strconv.Atoi(input)

	if age < 18 {
		fmt.Println("Your Age is too young")
	} else {
		fmt.Println("Your are eligible to vote")
	}

}
