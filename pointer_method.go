package main

import "fmt"

type Man struct {
	Name string
	Age  string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
	man.Age = "Umur " + man.Age
}

func main() {

	syahid := Man{"Syahid", "24"}

	syahid.Married()

	fmt.Println(syahid.Name, syahid.Age)
}
