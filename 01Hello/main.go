package main

import (
	"fmt"
	"strings"
)

const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

func main() {
	fmt.Println("Hello from AJS !")

	str := "Hello, 世界!"

	fmt.Println("Before : ", str)
	str = strings.ToLower(str)
	runes := []rune(str)

	var res string

	for val := range runes {
		res += string(runes[val])
	}
	fmt.Println("After : ", res)

}
