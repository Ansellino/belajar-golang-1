package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var d = 5
	var e = 2
	var c = a + b - d * e 
	fmt.Println(c);

	a += 10
	fmt.Println("Nilai a adalah", a, "Nilai b adalah", b)
}