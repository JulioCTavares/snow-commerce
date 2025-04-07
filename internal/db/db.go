package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/snowflakedb/gosnowflake"
)

var DB *sql.DB

func InitDB() {
	dsn := os.Getenv("URL_CONNECTION")

	var err error
	DB, err = sql.Open("snowflake", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to Snowflake: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to ping Snowflake: %v", err)
	}
}
