package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error

	godotenv.Load()

	host := os.Getenv("PG_HOST")
	port := "5432"
	user := os.Getenv("PG_USERNAME")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DB")

	fmt.Println("host", host)
	fmt.Println("user", user)
	fmt.Println("password", password)

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}
	fmt.Println("Connected to the database!")

	createTable()

}

func InsertServiceStatus(serviceName, serviceUrl, serviceStatus, elaspedTime string, requestTime time.Time) {

	query := `
	INSERT INTO services (service_name, service_url, status_code, elasped_time, request_time)
	VALUES ($1, $2, $3, $4, $5);
	`
	_, err := DB.Exec(query, serviceName, serviceUrl, serviceStatus, elaspedTime, requestTime)
	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
	}
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS services (
		id SERIAL PRIMARY KEY,
		service_name VARCHAR(100),
		service_url VARCHAR(100),
		status_code VARCHAR(5),
		elasped_time VARCHAR(100),
		request_time TIMESTAMP
	);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	fmt.Println("Table 'services' is ready.")
}
