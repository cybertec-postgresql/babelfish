package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/cybertec-postgresql/babelfish/babelfishdb"
)

var db *sql.DB

func main() {
	var err error

	db, err = babelfishdb.Open("sqlserver", "")
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	log.Printf("Connected!\n")
	// Close the database connection pool after program executes
	defer db.Close()

	SelectVersion()
}

// Gets and prints SQL Server version
func SelectVersion() {
	// Use background context
	ctx := context.Background()

	// Ping database to see if it's still alive.
	// Important for handling network issues and long queries.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	fmt.Printf("\nMSSQL version:\n%s\n", babelfishdb.Get(db, "SELECT @@version"))
	fmt.Printf("\nPgSQL version:\n%s\n", babelfishdb.Get(db, "SELECT version()"))

	fmt.Printf("\nMSSQL dbname:\n%s\n", babelfishdb.Get(db, "SELECT db_name()"))
	fmt.Printf("\nPgSQL dbname:\n%s\n", babelfishdb.Get(db, "SELECT current_database()"))
}
