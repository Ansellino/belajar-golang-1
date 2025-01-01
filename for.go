package main

import "fmt"

func main(){
	// Cara 1
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	fmt.Println("Selesai")

	// Cara 2
	for counterNum := 1; counterNum <= 10; counterNum++ {
		fmt.Println("Perulangan ke", counterNum)
	}
	fmt.Println("Selesai")

	// cara 3 Manual
	names := []string{"Jeremy","Ansellino","Gunawan"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// cara 4
	for index, name := range names{
		fmt.Println("Index", index, "=", name)
	}

	// cara 5
	for _, name := range names{
		fmt.Println(name)
	}
	
}