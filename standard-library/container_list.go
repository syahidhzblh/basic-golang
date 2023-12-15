package main

import (
	"container/list"
	"fmt"
)

func main() {

	//Container list biasanya digunakan untuk antrian
	data := list.New()
	data.PushBack("Syahid")
	data.PushBack("Hizbullah")
	data.PushBack("Ganteng")

	//Ambil data head atau data awal
	head := data.Front()

	for i := head; i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
