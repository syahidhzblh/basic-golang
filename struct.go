package main

import "fmt"

//Deklarasi Tipe Struct
type Customer struct {
	Name, Address string
	Age           int
}

//Struct Method

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {

	//Cara 1
	var syahid Customer

	syahid.Name = "Syahid"
	syahid.Address = "Tangerang"
	syahid.Age = 24

	fmt.Println(syahid)

	//Cara 2
	hisbul := Customer{
		Name:    "Hisbul",
		Address: "Jakarta",
		Age:     25,
	}

	fmt.Println(hisbul)

	//Cara 3
	said := Customer{"Said", "Tangerang", 20}
	fmt.Println(said)

	hisbul.sayHello("Syahid")
}
