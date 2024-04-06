package main

import (
	"log"
	"net"
	"strings"
)

func handleConn(conn net.Conn) {
	// Close connection after handling it
	defer conn.Close()

	// Read http connection
	buffer := make([]byte, 2048)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("ERROR: Could not read from connection [%s]", conn.RemoteAddr().String())
		conn.Close()
	}

	lines := strings.Split(string(buffer[:n]), "\n")

	// Parse http request line parameters to determine which file to serve
	httpRequestLineParams := strings.Split(lines[0], " ")
	for i, param := range httpRequestLineParams {
		log.Printf("Parameter %d: %s", i, param)
	}
}

func main() {
	// Listen and create a TCP server in localhost:port
	ln, err := net.Listen("tcp", "127.0.0.1:6969")
	if err != nil {
		log.Fatalf("An error occured while starting the server: %s", err)
	}

	// Accept every incoming connection and handle it properly concurrently
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("ERROR: Could not accept connection %s", conn.RemoteAddr().String())
		}
		log.Printf("Accepted connection from [%s]", conn.RemoteAddr().String())
		go handleConn(conn)
	}
}
