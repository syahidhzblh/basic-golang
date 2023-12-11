package main

import (
	"fmt"
)

func main() {

	const firstname string = "Syahid"
	const lastname = "Hizbullah"
	fmt.Println(firstname, lastname)

	const (
		nama   string = "Syahid Hizbullah"
		tempat        = "Jakarta"
	)

	fmt.Println(nama, tempat)
}
