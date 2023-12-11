package main

import "fmt"

func getFullName() (string, string) {

	return "Syahid", "Hizbullah"
}

func main() {

	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)

	//Ignore salah satu return value
	firstName, _ := getFullName()
	fmt.Println(firstName)
}
