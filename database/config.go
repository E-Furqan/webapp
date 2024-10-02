package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/E-Furqan/webapp/data"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	portStr := os.Getenv("port")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}
	var (
		user     = os.Getenv("user")
		password = os.Getenv("password")
		host     = os.Getenv("host")
		db_name  = os.Getenv("db_name")
	)
	var connectionstring = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, db_name)

	DB, err = gorm.Open(postgres.Open(connectionstring), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// Auto-migrate the schema
	err = DB.AutoMigrate(&data.Student_info{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connection established")
}
