package main

func getFullName() (firstName, middleName, lastName string) {

	firstName = "Syahid"
	middleName = "Hizbullah"
	lastName = "Ganteng"

	return firstName, middleName, lastName
}

func main() {

	a, b, c := getFullName()

	println(a, b, c)
}
