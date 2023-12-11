package main

import "fmt"

//Alias
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {

	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {

	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {

	spam := spamFilter

	sayHelloWithFilter("Syahid", spamFilter)
	sayHelloWithFilter("Anjing", spam)
}
