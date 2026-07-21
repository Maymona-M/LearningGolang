package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

var db *sql.DB

type PingResponse struct {
	Status string `json:"status"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := PingResponse{Status: "ok"}
	json.NewEncoder(w).Encode(response)
}

func ipsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT DISTINCT ip FROM readings")
	if err != nil {
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var ips []string
	for rows.Next() {
		var ip string
		if err := rows.Scan(&ip); err != nil {
			http.Error(w, "Failed to read IPs", http.StatusInternalServerError)
			return
		}
		ips = append(ips, ip)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ips)
}


func main() {
	var err error
	db, err = sql.Open("sqlite", "../../2_tcp_server/monitor.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/api/ping", pingHandler)
	http.HandleFunc("/api/ips", ipsHandler)

	fmt.Println("Server starting on :8081")
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
