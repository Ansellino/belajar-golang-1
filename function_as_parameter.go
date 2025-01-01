package main

import "fmt"

type filtered func(string) string

func sayHelloWithFilter(name string, filter filtered){
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string{
	if name == "Anjing"{
		return "..."
	}else {
		return name
	}
}

func main() {
	
	filter := spamFilter

	sayHelloWithFilter("Anjing", filter)
	sayHelloWithFilter("Jeremy", spamFilter)

}