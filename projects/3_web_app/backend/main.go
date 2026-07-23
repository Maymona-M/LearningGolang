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

type Reading struct {
	ID        int     `json:"id"`
	IP        string  `json:"ip"`
	CPU       float64 `json:"cpu"`
	Mem       float64 `json:"mem"`
	Disk      float64 `json:"disk"`
	Timestamp int64   `json:"timestamp"`
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

func readingsHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	if ip == "" {
		http.Error(w, "Missing 'ip' query parameter", http.StatusBadRequest)
		return
	}

	rows, err := db.Query("SELECT id, ip, cpu, mem, disk, timestamp FROM readings WHERE ip = ? ORDER BY timestamp DESC", ip)
	if err != nil {
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var readings []Reading
	for rows.Next() {
		var reading Reading
		if err := rows.Scan(&reading.ID, &reading.IP, &reading.CPU, &reading.Mem, &reading.Disk, &reading.Timestamp); err != nil {
			http.Error(w, "Failed to read readings", http.StatusInternalServerError)
			return
		}
		readings = append(readings, reading)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(readings)
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	var err error
	db, err = sql.Open("sqlite", "../../2_tcp_server/monitor.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	initUsersTable()
	createDefaultAdmin()

	// Routes Registered
	mux := http.NewServeMux()
	mux.HandleFunc("/api/ping", pingHandler)         // health check
	mux.HandleFunc("/api/ips", ipsHandler)           // returns list of unique sender IPs
	mux.HandleFunc("/api/readings", readingsHandler) // returns readings for a given ?ip=
	mux.HandleFunc("/api/login", loginHandler) // handles POST login requests, checks username/password

	fmt.Println("Server starting on :8081")
	err = http.ListenAndServe(":8081", enableCORS(mux))
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
