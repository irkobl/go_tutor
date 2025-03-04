package main

import (
	"bufio"
	"log"
	"net"
	"strings"
	"time"
)

const addr = "0.0.0.0:12345"

const proto = "tcp4"

func main() {
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	b, err := reader.ReadBytes('\n')
	if err != nil {
		log.Println(err)
		return
	}

	msg := strings.TrimSuffix(string(b), "\n")
	msg = strings.TrimSuffix(msg, "\r")

	if msg == "time" {
		conn.Write([]byte(time.Now().String() + "\n"))
	}

	conn.Close()
}
