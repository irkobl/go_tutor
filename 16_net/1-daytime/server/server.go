package main

import (
	"log"
	"net"
	"time"
)

func handler(conn net.Conn) {
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp4", ":3333")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		handler(conn)
	}
}
