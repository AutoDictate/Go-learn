package main

import "fmt"

const JWT = "7649182319" // Public

func main() {

	var username = "zoro"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T\n", username)

	var intValue int32 = 23
	fmt.Println(intValue)
	fmt.Printf("variable is of type : %T\n", intValue)

	var smallIntValue uint32 = 23
	fmt.Println(smallIntValue)
	fmt.Printf("variable is of type : %T\n", smallIntValue)

	var smallFloatValue float32 = 23
	fmt.Println(smallFloatValue)
	fmt.Printf("variable is of type : %T\n", smallFloatValue)

	// Implicit type

	var rollno = 89
	fmt.Println(rollno)
	fmt.Printf("variable is of type : %T\n", rollno)

	var name = "sanji"
	fmt.Println(name)
	fmt.Printf("variable is of type : %T\n", name)

	// no var style

	numberOfDaysInSept := 31
	fmt.Println("numberOfDaysInSept : ", numberOfDaysInSept)

	// constant values

	fmt.Println("JWT Token : ", JWT)

}
