package main

import "fmt"

type Persone struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Persone {
	return &Persone{
		Name: name,
		Age:  age,
	}
}

func main() {
	person := NewPerson("H", 14)
	fmt.Println(person)
}
