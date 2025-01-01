package main

import "fmt"

func main(){
	name := "Ansel"

	// variasi 1
	switch name {
	case "Ansel":
		fmt.Println("Hello Ansel")
	case "Budi":
		fmt.Println("Hello Budi")
	case "Jono":
		fmt.Println("Hello Jono")
	default:
		fmt.Println("Hi, boleh kenalan?")
	}

	// variasi 2
	switch length := len(name); length > 5{
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// variasi 3
	name = "Jeremy"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}