package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PingResponse struct {
	Status string `json:"status"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := PingResponse{Status: "ok"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api/ping", pingHandler)

	fmt.Println("Server starting on :8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
