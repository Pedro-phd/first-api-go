package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// Loading .env file
	errenv := godotenv.Load()
	if errenv != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Printf(" 🚀 System online in %s ambient", os.Getenv("AMBIENT"))
}
