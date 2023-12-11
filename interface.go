package main

import "fmt"

type HasName interface {
	getName() string
}

func sayHello(value HasName) {
	fmt.Println("Hello", value.getName())
}

type Person struct {
	Name string
}

type Soulmate struct {
	Name, Hobby string
	Age         int
}

func (person Person) getName() string {
	return person.Name
}

func (soulmate Soulmate) getName() string {
	return
}

func main() {

	person := Person{"Syahid"}
	sayHello(person)

	soulmate := Soulmate{
		Name:  "Zahwa",
		Hobby: "Running",
		Age:   25,
	}
	sayHello(soulmate)
}
