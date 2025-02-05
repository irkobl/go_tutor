package main

import "fmt"

type Animals interface {
	Voice()
}

type Dog struct {
	Name string
}

func (d *Dog) Voice() {
	fmt.Printf("%s: Woof\n", d.Name)
}

type Cat struct {
	Name string
}

func (c *Cat) Voice() {
	fmt.Printf("%s: Meow\n", c.Name)
}

func MakeVoice(a Animals) {
	a.Voice()
}

func main() {
	animals := []Animals{
		&Dog{Name: "sharic"},
		&Cat{Name: "murka"},
	}

	for _, anim := range animals {
		MakeVoice(anim)
	}
}
