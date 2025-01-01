package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	
	for _, number := range numbers {
		total += number
	}

	return total
}
func main() {
	fmt.Println(sumAll(10,10,10))
	fmt.Println(sumAll(20,20,20,20,20))
	fmt.Println(sumAll(15,15))

	//slice number
	numbers := []int{10,10,10,10}
	fmt.Println(sumAll(numbers...))
}