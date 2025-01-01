package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main(){
	result := getHello("Ansel")
	fmt.Println(result)

	fmt.Println(getHello("Budi"))
	fmt.Println(getHello("Joni"))
}