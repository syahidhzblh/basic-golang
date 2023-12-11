package main

import "fmt"

func main() {

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	fmt.Println(days)

	slice1 := days[5:]
	slice1[0] = "Sabtu Baru"
	slice1[1] = "Minggu Baru"

	fmt.Println(slice1)
	fmt.Println(days)

	//Append Session

	slice2 := append(slice1, "Libur Baru")
	slice2[0] = "Sabtu Lama"
	fmt.Println(slice2)

	fmt.Println(days)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	//Perbedaan Array dan Slice

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
