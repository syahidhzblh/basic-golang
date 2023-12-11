package main

import "fmt"

func getGoodBye(name string) string {
	return "Good bye " + name
}

func main() {

	goodbye := getGoodBye

	fmt.Println(goodbye("Syahid"))
}
