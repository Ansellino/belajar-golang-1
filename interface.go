package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

// person
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}
// animal
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Ansel"}
	sayHello(person)

	animal := Animal{Name: "Kucing"}
	sayHello(animal)
}
