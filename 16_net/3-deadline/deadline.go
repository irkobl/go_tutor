package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func handler(conn net.Conn) {
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(time.Second * 5))
	r := bufio.NewReader(conn)
	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}
		fmt.Println(string(msg))
		_, err = conn.Write(msg)
		if err != nil {
			return
		}
		conn.SetDeadline(time.Now().Add(time.Second * 5))
	}

}

func main() {
	listener, err := net.Listen("tcp4", ":12345")
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
