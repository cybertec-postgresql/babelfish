package babelfishdb

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

func Open(protocol string, conn string) (db *sql.DB, err error) {

	// Create connection string
	if conn == "" {
		conn = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", server, user, password, port)
	}
	// Create protocol string
	if protocol == "" {
		protocol = "sqlserver"
	}
	// Create connection pool
	db, err = sql.Open(protocol, conn)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	return
}

func Get(db *sql.DB, sql string) (result string) {
	err := db.QueryRowContext(context.Background(), sql).Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	return
}
