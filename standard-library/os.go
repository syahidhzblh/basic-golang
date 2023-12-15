package main

import (
	"fmt"
	"os"
)

func main() {

	s := os.Args
	for _, isi := range s {
		fmt.Println(isi)
	}

	hasil, err := os.Hostname()
	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println(err)
	}
}
