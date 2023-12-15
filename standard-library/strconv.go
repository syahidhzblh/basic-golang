package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Konversi from String to Bool
	hasil, err := strconv.ParseBool("true")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(hasil)
	}

	//Konversi dari String to Int dengan "Atoi"
	hasilInt, err := strconv.Atoi("1000")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(hasilInt)
	}

	//Konversi dari Int to String dengan "Itoa"
	hasilString := strconv.Itoa(999)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(hasilString)
	}
}
