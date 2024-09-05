package main

import (
	"fmt"
	"math/rand"
)

var i int

func main() {

	fmt.Println("Let's Create a Dice Game using Switch case")

	// rand.Seed is deprecated after -> go 1.20 version

	var isBool = true

	for isBool {
		diceNumber := rand.Intn(6) + 1
		i++
		switch diceNumber {
		case 1:
			fmt.Println("dice number is 1")
			fallthrough
		case 2:
			fmt.Println("dice number is 2")
			fallthrough
		case 3:
			fmt.Println("dice number is 3")
			fallthrough
		case 4:
			fmt.Println("dice number is 4")
			fallthrough
		case 5:
			fmt.Println("dice number is 5")
			fallthrough
		case 6:
			fmt.Println("dice number is 6")
			isBool = false
			break
		default:
			fmt.Println("dice number is default")
		}
	}

	fmt.Println(i)
}
