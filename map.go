package main

import "fmt"

func main() {
	/* var person = map[string]string{}
	person["name"] = "Ansel"
	person["address"] = "Jakarta" */

	person := map[string]string{
		"name" : "Jeremy",
		"address" : "Jakarta",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)
}