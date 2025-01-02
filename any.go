package main

import "fmt"

func Ups() any {
	return 1
}

func main() {
	var kosong = Ups()
	fmt.Println(kosong)
}