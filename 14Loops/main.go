package main

import (
	"fmt"
)

func main() {

	fmt.Println("Let's learn some loops and breaks")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// Method 1 -> Classic Method
	fmt.Println("------------------------- Classic Method ---------------------------")
	for i := 0; i < len(days); i++ {
		fmt.Println(i, days[i])
	}

	// Method 2 -> Only get index to get the array value
	fmt.Println("------------------------- Index value only ---------------------------")
	for i := range days {
		fmt.Println(days[i])
	}

	// Method 3 -> Enhanced For loop
	fmt.Println("------------------------- Actual value ---------------------------")
	for index, day := range days {
		fmt.Println(index, day)
	}

	fmt.Println("------------------------- While Loop using for ---------------------------")

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 9 {
			rogueValue++
			goto message
		}

		fmt.Println(rogueValue)
		rogueValue++
	}

message:
	fmt.Println("Vanakam da Mapla, Label la irundhu... !!!!")

}
