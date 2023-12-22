package main

import "fmt"

// Start of Kontrak Interface HasName
type HasName interface {
	getName() string
	getHobby() string
}

func sayHello(value HasName) {
	fmt.Println("Hello", value.getName())
	fmt.Println("My Hobby is", value.getHobby())
} // End of Kontrak Interface HasName

// Implementasi
type Person struct {
	Name, Hobby string
}

func (person Person) getName() string {
	return person.Name
}

func (person Person) getHobby() string {
	return person.Hobby
}

type Soulmate struct {
	Name, Hobby string
}

func (soulmate Soulmate) getName() string {
	return soulmate.Name
}

func (soulmate Soulmate) getHobby() string {
	return soulmate.Hobby
}

func main() {

	person := Person{
		Name:  "Syahid",
		Hobby: "Playing Game",
	}
	sayHello(person)

	soulmate := Soulmate{
		Name:  "Zahwa",
		Hobby: "Running",
	}
	sayHello(soulmate)
}
