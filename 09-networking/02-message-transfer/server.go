package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Server started...")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Waiting for connection...")

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Client connected!")

	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reaing:", err)
		return
	}

	fmt.Println("Message received:", message)
}