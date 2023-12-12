package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "test"
	dbname   = "feast"
)

func main() {
	// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")

	// Create a table if it does not exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50),
			email VARCHAR(100)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table 'users' created or already exists")

	// Insert sample data into the table
	_, err = db.Exec(`
		INSERT INTO users (username, email) VALUES
			('user1', 'user1@example.com'),
			('user2', 'user2@example.com')
	`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sample data inserted into the 'users' table")
}
