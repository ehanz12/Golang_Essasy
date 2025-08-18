package database

import "fmt"

var connection string

func init() {
	fmt.Println("Initializing database connection...")
	connection = "localhost:5432"
}

func GetConnection() string {
	return connection
}