package main

import (
	"fmt"
	"regexp"
)

func main() {

	//Find Upper case
	regex := regexp.MustCompile(`^[[:lower:]]+$`)

	fmt.Println(regex.MatchString("syahid"))
	fmt.Println(regex.MatchString("syUhid"))
	fmt.Println(regex.MatchString("sywhid"))
}
