package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// creates the users table if it doesn't already exist
func initUsersTable() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		role TEXT NOT NULL
	);` // schema: id, username, hashed password, role

	_, err := db.Exec(createUsersTable)
	if err != nil {
		log.Fatal(err)
	}
}

// creates a default admin account, only if one doesn't already exist
func createDefaultAdmin() {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE role = 'admin'").Scan(&count) // check if an admin already exists
	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		return // admin already exists, don't create another
	}

	defaultPassword := "admin123" // plaintext, only used once, right here

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost) // hash before storing
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		"INSERT INTO users (username, password_hash, role) VALUES (?, ?, ?)",
		"admin", string(hashedPassword), "admin",
	) // insert the new admin row
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Default admin account created (username: admin)") // confirmation message
}

// shape of the incoming login request body
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// shape of what we send back after a login attempt
type LoginResponse struct {
	Success bool   `json:"success"`
	Role    string `json:"role,omitempty"`
	Message string `json:"message,omitempty"`
}

// shape of the incoming "create user" request body
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// shape of a user record returned to the frontend (no password hash included)
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) // only DELETE allowed
		return
	}

	idStr := r.URL.Query().Get("id") // read ?id= from the URL
	if idStr == "" {
		http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("DELETE FROM users WHERE id = ?", idStr) // delete the matching row
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted"}) // confirm success
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, username, role FROM users") // only select the 3 safe columns
	if err != nil {
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.Role); err != nil { // read each row into a User
			http.Error(w, "Error reading row", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users) // send back the full list as JSON
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) // only POST allowed
		return
	}

	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { // parse incoming JSON
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" || req.Role == "" { // basic validation
		http.Error(w, "Username, password, and role are required", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost) // hash before storing
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec(
		"INSERT INTO users (username, password_hash, role) VALUES (?, ?, ?)",
		req.Username, string(hashedPassword), req.Role,
	) // insert the new user
	if err != nil {
		http.Error(w, "Username already exists", http.StatusConflict) // UNIQUE constraint failed
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"}) // confirm success
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) // only POST allowed here
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { // parse incoming JSON into req
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var storedHash, role string
	err := db.QueryRow(
		"SELECT password_hash, role FROM users WHERE username = ?", req.Username,
	).Scan(&storedHash, &role) // look up the user by username

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(LoginResponse{Success: false, Message: "Invalid username or password"}) // vague on purpose
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(req.Password)) // check entered password against stored hash
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(LoginResponse{Success: false, Message: "Invalid username or password"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{Success: true, Role: role}) // login succeeded
}
