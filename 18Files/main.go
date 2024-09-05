package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Let's learn how files work in Go lang")

	content := "Hello from backend developer of JAVA"

	file, err := os.Create("./surya.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("Length of the file", length)
	_ = file.Close()
}
