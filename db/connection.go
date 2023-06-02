package db

import (
	"database/sql"
	"fmt"

	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() error {
	godotenv.Load()
	connStr := os.Getenv("DB_CONNECTION_STRING")
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	DB = db
	fmt.Println("Connected to the database!")

	return nil
}
