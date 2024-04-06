package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:6969")
	if err != nil {
		log.Fatalf("An error occured while starting the server: %s", err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("ERROR: Could not accept connection %s", conn.RemoteAddr().String())
		}
		log.Printf("Accepted connection from [%s]", conn.RemoteAddr().String())
		buffer := make([]byte, 2048)
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("ERROR: Could not read from connection [%s]", conn.RemoteAddr().String())
		}

		log.Printf("Connection body: \n %s", buffer)
		log.Printf("Connection length %d", n)
	}
}
