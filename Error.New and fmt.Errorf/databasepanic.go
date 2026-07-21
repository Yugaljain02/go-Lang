package main

import (
	"fmt"
	"os"
)

func connectDatabase() {
	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		panic("DB_URL is not configured")
	}

	fmt.Println("Connected to database:", dbURL)
}

func main() {
	fmt.Println("Application Starting...")

	connectDatabase()

	fmt.Println("Application Running...")
}