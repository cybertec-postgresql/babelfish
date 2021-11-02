package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

// Replace with your own connection parameters
// Ok, you now have 2 babelfish:
// Single-db: PostgreSQL port 5433, TDS port 1433
// Multi-db: PostgreSQL port 5434, TDS port 1434
// Shared data: User, Database, Password = hs
var server = "10.1.136.69"
var port = 1433
var user = "hs"
var password = "hs"

var db *sql.DB

func main() {
	var err error

	// Create connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d",
		server, user, password, port)

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
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

	var result string

	// Run query and scan for result
	err = db.QueryRowContext(ctx, "SELECT @@version").Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("\n%s\n", result)
}
