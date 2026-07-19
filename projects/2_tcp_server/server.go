package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"

	"learninggolang/projects/2_tcp_server/stats"
)


func handleConnection(conn net.Conn) {
	defer conn.Close()

	remoteAddr := conn.RemoteAddr().String()
	ip, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		ip = remoteAddr
	}

	fmt.Println("Client connected from:", ip)

	var payload stats.SystemStats
	decoder := gob.NewDecoder(conn)
	if err := decoder.Decode(&payload); err != nil {
		fmt.Println("Error decoding stats:", err)
		return
	}

	fmt.Printf("From %s -> CPU: %.1f%% | Mem: %.1f%% | Disk: %.1f%% | Time: %d\n",
		ip, payload.CPU, payload.Mem, payload.Disk, payload.Time)

	fmt.Fprintln(conn, "ACK: message received")
}

func main() {
	fmt.Println("Server started...")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Waiting for connections...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}