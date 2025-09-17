package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

var db *sql.DB

func initDB() {
	host := getEnv("DB_HOST", "localhost")
	name := getEnv("DB_NAME", "booksdb")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "password")
	port := getEnv("DB_PORT", "5432")

	conStr := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s port=%s sslmode=disable",
		host, name, user, password, port,
	)

	var err error
	db, err = sql.Open("postgres", conStr)
	if err != nil {
		log.Fatal("Failed to open Database:", err)
	}

	// ตรวจสอบว่าเชื่อมต่อได้จริง
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to Database:", err)
	}

	log.Println("Connected to Database successfully!")
}

func main() {
	initDB()
}
