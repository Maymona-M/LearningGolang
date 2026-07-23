package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost) // hash it before storing
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
