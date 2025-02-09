package main

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"os"
)

type serializer interface {
	serialize() ([]byte, error)
}

type person struct {
	Name string
	Age  int
}

type car struct {
	Model string
	Year  int
}

func (p *person) serialize() ([]byte, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (c *car) serialize() ([]byte, error) {
	b, err := xml.Marshal(c)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func store(s serializer, wc io.WriteCloser) error {
	b, err := s.serialize()
	if err != nil {
		return err
	}
	_, err = wc.Write(b)
	return err
}

func main() {
	p := person{
		Name: "Kurt",
		Age:  27,
	}

	var s serializer = &p

	f, err := os.Create("./output.json")
	if err != nil {
		log.Fatal(err)
	}

	err = store(s, f)
	if err != nil {
		log.Fatal(err)
	}

	c := car{
		Model: "Beet",
		Year:  1970,
	}

	err = store(&c, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

}
