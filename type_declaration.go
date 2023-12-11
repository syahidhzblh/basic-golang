package main

import "fmt"

func main() {
	type NoKTP string

	var KTP1 NoKTP = "31313131"

	var KTP2 = "21212121"
	var ContohKTP NoKTP = NoKTP(KTP2)

	fmt.Println(KTP1)
	fmt.Println(KTP2)
}
