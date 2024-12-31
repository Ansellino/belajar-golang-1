package main

import "fmt"

func main() {

	type NoKTP string

	var ktpAnsel NoKTP = "11111"

	var contoh = "22222"
	var contohKtp = NoKTP(contoh)

	fmt.Println(ktpAnsel)
	fmt.Println(contohKtp)
}