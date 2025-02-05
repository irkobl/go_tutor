package main

import "fmt"

type Engine struct {
	Model string
}

type Car struct {
	Make  string
	Color string
	Engine
}

func main() {
	car := Car{Make: "BMW", Color: "bl", Engine: Engine{"s63"}}

	fmt.Println(car.Make, car.Color)
	fmt.Println(car.Model)

}
