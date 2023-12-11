package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	address1 := Address{"Tangerang", "Banten", "Indonesia"}
	address2 := &address1

	address2.City = "Serang"

	fmt.Println(address1)
	fmt.Println(address2)

	//Asterisk Operator

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
}
