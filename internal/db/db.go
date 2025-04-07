package db

import (
	"database/sql"
	"log"

	_ "github.com/snowflakedb/gosnowflake"
)

var DB *sql.DB

func InitDB() {
	dsn := "JULIOCTAVARES:tokkat-jubbo8-sYscej@KMRHBCK-NA73918/PROJECT_DB/PUBLIC?warehouse=COMPUTE_WH"

	var err error
	DB, err = sql.Open("snowflake", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to Snowflake: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to ping Snowflake: %v", err)
	}
}
