package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DataBaseConectionString is the MySQL connection string
	DataBaseConectionString = ""

	// Port where the API will be running
	Port = 0
)

// Load will initialize environment variables
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Error initializing environment variables", err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	DataBaseConectionString = fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}
