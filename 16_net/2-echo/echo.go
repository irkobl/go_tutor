package main

import (
	"bufio"
	"log"
	"net"
)

func handler(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}
		msg = append(msg, '\n')
		_, err = conn.Write(msg)
		if err != nil {
			return
		}
	}
}

func main() {
	listen, err := net.Listen("tcp4", ":12345")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handler(conn)
	}
}
