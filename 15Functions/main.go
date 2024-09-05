package main

import (
	"fmt"
	"strings"
)

func main() {

	res, _ := IsPalindrome("Surya")
	fmt.Println(res)

	fmt.Println(proAdder(1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 6))
}

func IsPalindrome(name string) (bool, error) {

	name = strings.ToLower(name)
	for i := 0; i < len(name); i++ {
		if name[i] != name[len(name)-i-1] {
			return false, nil
		}
	}
	return true, nil
}

func proAdder(values ...int) int {

	total := 0
	for _, v := range values {
		total ^= v
	}
	return total
}
