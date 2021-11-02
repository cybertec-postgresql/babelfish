package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/cybertec-postgresql/babelfish/babelfishdb"
)

var db *sql.DB

const ddl = `
CREATE SCHEMA TestSchema;

CREATE TABLE TestSchema.Employees (
  Id       INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
  Name     NVARCHAR(50),
  Location NVARCHAR(50)
);

INSERT INTO TestSchema.Employees (Name, Location) VALUES
  (N'Jared',  N'Australia'),
  (N'Nikita', N'India'),
  (N'Tom',    N'Germany');
`

func main() {
	var err error

	db, err = babelfishdb.Open("sqlserver", "")
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	// Ping database to see if it's still alive.
	// Important for handling network issues and long queries.
	err = db.PingContext(context.Background())
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	log.Printf("Connected!\n")
	// Close the database connection pool after program executes
	defer func() {
		_, err = db.Exec("DROP TABLE TestSchema.Employees; DROP SCHEMA TestSchema;")
		if err != nil {
			log.Print("Cannot drop TestSchema: " + err.Error())
		}
		db.Close()
	}()

	_, err = db.Exec(ddl)
	if err != nil {
		log.Print("Cannot create TestSchema: " + err.Error())
	}

	rand.Seed(time.Now().UnixNano())
	SelectEmployee(rand.Intn(3) + 1)
}

// Gets and prints SQL Server version
func SelectEmployee(id int) {
	fmt.Printf("\nEmployee #%d:\n%s\n", id, babelfishdb.Get(db, fmt.Sprintf(`
SELECT [Name] + ' from ' + [Location]
FROM TestSchema.Employees WHERE Id = %d`, id)))
}
