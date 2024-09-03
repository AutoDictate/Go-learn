package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	fmt.Println("Welcome to time study of go lang")

	t := time.Date(2003, 06, 11, 12, 12, 12, 0, time.Now().Location())

	cpu := runtime.NumCPU()
	fmt.Println("Number of cpu is running : ", cpu)

	fmt.Println("the current time is", t.Format("2006-01-02 15:04:05"))
}
