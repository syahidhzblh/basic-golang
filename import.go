package main

import (
	"basic-golang/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Syahid")
	fmt.Println(result)

	fmt.Println(helper.SayGoodbye("Syahid"))
	fmt.Println(helper.Application)
}
