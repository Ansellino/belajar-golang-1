package main

import "fmt"

func main(){
	var names [3] string

	names[0] = "Jeremy"
	names[1] = "Ansellino"
	names[2] = "Gunawan"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int {
		90,
		15,
	}

	var titikTiga = [...]int{
		100,
		40,
		20,
		10,
	}

	fmt.Println(titikTiga)
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
}