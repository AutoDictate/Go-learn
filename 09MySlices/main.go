package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("--------------------- Initializing Slice ------------------------")

	var nameList = []string{"Surya", "Moha", "Selva"}

	nameList = append(nameList, "Sibi", "Karthi", "Vimal", "Sakthi")

	nameList = append(nameList[:3])

	fmt.Println(nameList)

	highScores := make([]int, 4)

	highScores[0] = 23
	highScores[1] = 25
	highScores[2] = 28
	highScores[3] = 29
	//highScores[4] = 27

	highScores = append(highScores, 27, 28, 29, 30)

	fmt.Println(highScores)
	fmt.Println("Length : ", len(highScores))
	fmt.Println("Capacity : ", cap(highScores))

	// Sort Package
	fmt.Println("--------------------- Sorting Slice ------------------------")

	fmt.Println(highScores)
	fmt.Println("Before sorted : ", sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("After sorted : ", sort.IntsAreSorted(highScores))
	fmt.Println(highScores)

	// How to remove a value from the slices based on index

	fmt.Println("--------------------- Remove slice ------------------------")

	var namesList = []string{"surya", "selva", "moha", "sibi", "ezhil", "karthi", "vimal", "sakthi", "chandru"}

	fmt.Println("Initial name list", namesList)
	fmt.Println("Total numbers : ", len(namesList))

	newNames := append(namesList[:4], namesList[5:]...)

	fmt.Println("Final name list : ", newNames)
	fmt.Println("Total numbers : ", len(newNames))

	fmt.Println("No of persons removed : ", len(namesList)-len(newNames))

}
