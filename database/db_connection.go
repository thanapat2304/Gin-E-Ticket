package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB // We will simulate this for mock data

func InitDB() {
	// Simulate a successful database connection for portfolio
	log.Println("Successfully connected to mock database for portfolio showcase")
	// In a real scenario, you would initialize your mock DB here
	// For example, if you had a mock implementation of sql.DB methods, you would assign it here.
	// As this is a portfolio piece, we'll just indicate a successful mock connection.
}
