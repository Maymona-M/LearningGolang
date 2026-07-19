package main

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"time"

	_ "modernc.org/sqlite"

	"learninggolang/projects/2_tcp_server/stats"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite", "./monitor.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS readings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ip TEXT,
		cpu REAL,
		mem REAL,
		disk REAL,
		timestamp INTEGER
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}

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

	readableTime := time.Unix(payload.Time, 0).Format("03:04:05 PM")

	fmt.Printf("From %s -> CPU: %.1f%% | Mem: %.1f%% | Disk: %.1f%% | Time: %s\n",
		ip, payload.CPU, payload.Mem, payload.Disk, readableTime)

	_, err = db.Exec(
		"INSERT INTO readings (ip, cpu, mem, disk, timestamp) VALUES (?, ?, ?, ?, ?)",
		ip, payload.CPU, payload.Mem, payload.Disk, payload.Time,
	)
	if err != nil {
		fmt.Println("Error saving to database:", err)
	}

	fmt.Fprintln(conn, "ACK: message received")
}

func main() {
	initDB()
	defer db.Close()

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