package main

import (
	"fmt"
	"log"
)

type modern interface {
	isModern() bool
}

type old interface {
	isOld() bool
	isObsolete() bool
}

type cassete struct {
	label string
}

func (c cassete) isModern() bool {
	return false
}

func (c cassete) isOld() bool {
	return true
}

func (c *cassete) isObsolete() bool {
	return false
}

func main() {
	var m modern
	var c cassete

	m = c

	// m.label = "Ace of Base" - ошибка
	m.isModern()

	switch c1 := m.(type) {
	case cassete:
		c1.label = "Ace of Base"
	case old:
		c1.isOld()
	}

	if c1, ok := m.(cassete); ok {
		c1.label = "Ace of base"
		log.Printf("%+v", c1)
	}

	if c2, ok := m.(old); ok {
		fmt.Println("Old: ", c2.isOld(), "\tObsolete: ", c2.isObsolete())
	}
}
