package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name   string
	Age    int
	status bool
}

func (u User) GetStatus() bool {
	fmt.Println("The status of the user is :", u.status)
	return u.status
}

func (u User) GetEmail() string {
	newEmail := strings.ToLower(u.Name) + "" + strconv.Itoa(u.Age) + "@gmail.com"
	return newEmail
}

func main() {

	fmt.Println("Let's learn some Methods concept")

	surya := User{
		Name:   "Surya",
		Age:    22,
		status: true,
	}

	fmt.Println(surya.GetStatus())
	fmt.Println(surya.GetEmail())
}
