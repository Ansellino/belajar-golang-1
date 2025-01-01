package main

import "fmt"

func getFullName() (string, string) {
	return "Jeremy", "Ansellino"
}

func main(){
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
	namaDepan, _ := getFullName()
	fmt.Println(namaDepan)
}