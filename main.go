package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB
var err error

// Replace with your own connection parameters
// Ok, you now have 2 babelfish:
// Single-db: PostgreSQL port 5433, TDS port 1433
// Multi-db: PostgreSQL port 5434, TDS port 1434
// Shared data: User, Database, Password = hs
const (
	server   = "10.1.136.69"
	port     = 1433
	user     = "hs"
	password = "hs"
)

const create_ddl = `
CREATE SCHEMA TestSchema;

CREATE TABLE TestSchema.Employees (
	Id       INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
	Name     NVARCHAR(50),
	Location NVARCHAR(50)
	);
	
	INSERT INTO TestSchema.Employees (Name, Location) VALUES
	(N'Hans',   N'Austria'),
	(N'Pavlo',  N'Ukraine'),
	(N'Ants',   N'Estonia');
`

const drop_dll = `
IF OBJECT_ID('TestSchema.Employees', 'U') IS NOT NULL DROP TABLE TestSchema.Employees; 
DROP SCHEMA TestSchema;
`

const select_sql = `
DECLARE @MyID INT;
SET @MyID = @ID;
SELECT [Name] + ' from ' + [Location] FROM TestSchema.Employees WHERE Id = @MyID;
`

func main() {
	conn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", server, user, password, port)
	// Create connection pool
	if db, err = sql.Open("sqlserver", conn); err != nil {
		return
	}
	// Ping database to establish a real connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}

	log.Printf("Connected!\n")

	// Close the database connection pool after program executes
	defer db.Close()

	// Get and print SQL Server version and database name
	var (
		ver  string
		name string
	)
	err = db.QueryRow("SELECT @@version, db_name()").Scan(&ver, &name)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("\nMSSQL version:\n%s\n", ver)
	fmt.Printf("\nMSSQL dbname:\n%s\n", name)

	// Create test schema
	_, err = db.Exec(create_ddl)
	if err != nil {
		log.Print("Cannot create TestSchema: " + err.Error())
	}

	// Delete test database schema at the end of the session
	defer func() {
		_, err = db.Exec(drop_dll)
		if err != nil {
			log.Print("Cannot drop TestSchema: " + err.Error())
		}
	}()

	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(3) + 1
	var result string
	err = db.QueryRow(select_sql, sql.Named("ID", id)).Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("\nEmployee #%d:\n%s\n", id, result)
}
