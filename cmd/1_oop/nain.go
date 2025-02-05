package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

type Car struct {
	Make  string
	Color string
}

type Driver struct {
	Name string
	Car  Car
}

func main() {
	person := NewPerson("Ivan", 14)
	fmt.Println(person)

	dr := Driver{Name: "igor", Car: Car{Make: "BMV", Color: "blue"}}
	fmt.Println(dr.Name)
	fmt.Println(dr.Car.Make)
	// fmt.Println(dr.Color)

}
