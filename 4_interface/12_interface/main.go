package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	if r.Height > 0 && r.Width > 0 {
		return r.Height * r.Width
	}
	return -1
}

func GetArea(s Shape) float64 {
	return s.Area()
}

func main() {
	rectangle := Rectangle{5, 6}
	fmt.Println(GetArea(rectangle))
}
