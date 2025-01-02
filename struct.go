package main

import "fmt"

type Customer struct {
	Name, Address string
	Age				int
}

func main() {
	var ansel Customer
	fmt.Println("Ini sebelum di masukin structnya", ansel)

	ansel.Name = "Jeremy Ansellino Gunawan"
	ansel.Address = "Indonesia"
	ansel.Age = 24

	fmt.Println("ini sesudah", ansel)
	fmt.Println("Ini nama", ansel.Name)
	fmt.Println("Ini alamat", ansel.Address)
	fmt.Println("Ini umur", ansel.Age)

	joko := Customer{
		Name: "Joko",
		Address: "Indonesia",
		Age: 30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi","Indonesia", 30}
	fmt.Println(budi)
}