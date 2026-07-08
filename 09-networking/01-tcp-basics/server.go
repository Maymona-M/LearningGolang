package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Server is starting...")

	listener, err := net.Listen("tcp", ":8080") // opens listener, creates port at 8080
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Waiting for connection")

	conn, err := listener.Accept() // wait until one client connects to port
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Client connected")
}
