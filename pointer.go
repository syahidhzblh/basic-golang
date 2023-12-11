package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	address1 := Address{"Tangerang", "Banten", "Indonesia"}
	address2 := &address1 // Gunakan operator and & untuk membuat pointer

	address2.City = "Jakarta"

	fmt.Println(address1) //Data address 1 akan berubah mengikuti data address 2 karena menggunakan pointer
	fmt.Println(address2)
}
