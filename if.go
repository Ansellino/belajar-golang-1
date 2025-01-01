package main

import "fmt"

func main(){
	name := "Ansel"

	if name == "Ansel" {
		fmt.Println("Hello Ansel")
	} else if name == "Joki" {
		fmt.Println("Hello Joki")
	} else {
		fmt.Println("Hi, Boleh kenalan?")
	}

	// short statement
	if length := len(name); length > 5{
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}