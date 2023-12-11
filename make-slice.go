package main

import "fmt"

func main() {

	newSlice := make([]string, 2, 5)
	newSlice[0] = "syahid"
	newSlice[1] = "hizbullah"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Ganteng")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))
}
