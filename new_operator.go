package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	address1 := new(Address) //Operator new bisa untuk pengganti pointer, tapi hanya bisa untuk data awal yang kosong
	address2 := address1

	address2.City = "Tangerang"

	fmt.Println(address1)
	fmt.Println(address2)

}
