package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp4", "localhost:3333")
	if err != nil {
		log.Fatal(err)
	}
	msg, err := io.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ответ сервера: ", string(msg))
}
