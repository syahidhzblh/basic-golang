package main

import "fmt"

type Person struct {
	Name, Location, Hobby string
	Age                   int
}

func main() {

	person := Person{
		Name:     "Syahid",
		Location: "Tangerang",
		Hobby:    "Running",
		Age:      24,
	}

	fmt.Println(person)
	fmt.Printf("My name is %v \nI lives in %v \nMy hobby is %v \nI am %o Years Old", person.Name, person.Location, person.Hobby, person.Age)
}
