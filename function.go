package main

import "fmt"

func sayHello() {

	fmt.Println("Hello")
}

//Function with Parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

//Main Function for calling or execute other function
func main() {
	sayHello()
	sayHelloTo("Syahid", "Hizbullah")
}
